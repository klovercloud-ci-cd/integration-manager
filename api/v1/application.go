package v1

import (
	"github.com/klovercloud-ci-cd/integration-manager/api/common"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/api"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"github.com/labstack/echo/v4"
	"log"
	"strconv"
	"strings"
)

type applicationApi struct {
	companyService service.Company
	observerList   []service.Observer
}

// GetAll.. Get All Applications
// @Summary Get All Applications
// @Description Get All Applications
// @Tags Application
// @Produce json
// @Param companyId query string true "company id"
// @Success 200 {object} common.ResponseDTO{data=[]v1.Application}
// @Router /api/v1/applications [GET]

func (a applicationApi) GetAll(context echo.Context) error {
	companyId := context.QueryParam("companyId")
	if companyId == "" {
		return context.JSON(404, common.ResponseDTO{
			Message: "company id is required",
		})
	}
	option := getQueryOption(context)
	option.LoadRepositories = true
	option.LoadApplications = true
	data, total := a.companyService.GetAllApplications(companyId, option)
	metadata := common.GetPaginationMetadata(option.Pagination.Page, option.Pagination.Limit, total, int64(len(data)))
	uri := strings.Split(context.Request().RequestURI, "?")[0]
	if option.Pagination.Page > 0 {
		metadata.Links = append(metadata.Links, map[string]string{"prev": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page-1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	metadata.Links = append(metadata.Links, map[string]string{"self": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	if (option.Pagination.Page+1)*option.Pagination.Limit < metadata.TotalCount {
		metadata.Links = append(metadata.Links, map[string]string{"next": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page+1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	return common.GenerateSuccessResponse(context, data,
		&metadata, "Successful")
}

// Get.. Get Application by Application id
// @Summary Get Application by Application id
// @Description Gets Application by Application id
// @Tags Application
// @Produce json
// @Param id path string true "Application id"
// @Param companyId query string true "company id"
// @Param repositoryId query string true "repository id"
// @Success 200 {object} common.ResponseDTO{data=v1.Application}
// @Router /api/v1/applications/{id} [GET]
func (a applicationApi) GetById(context echo.Context) error {
	id := context.Param("id")
	if id == "" {
		return common.GenerateErrorResponse(context, nil, "application Id is required!")
	}
	companyId := context.QueryParam("companyId")
	if companyId == "" {
		return context.JSON(404, common.ResponseDTO{
			Message: "company id is required",
		})
	}
	repositoryId := context.QueryParam("repositoryId")
	if repositoryId == "" {
		return context.JSON(404, common.ResponseDTO{
			Message: "repository id is required",
		})
	}
	data := a.companyService.GetApplicationByApplicationId(companyId, repositoryId, id)
	if data.MetaData.Id == "" {
		return common.GenerateErrorResponse(context, nil, "Company not found!")
	}
	return common.GenerateSuccessResponse(context, data,
		nil, "Successful")
}

// Update... Update Application
// @Summary  Update Application
// @Description Update Application by company id and  repository id
// @Tags Application
// @Accept json
// @Produce json
// @Param data body v1.ApplicationsDto true "ApplicationsDto Data"
// @Param companyId query string true "company id"
// @Param repositoryId query string true "repository id"
// @Success 200 {object} common.ResponseDTO
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/applications [POST]
func (a applicationApi) Update(context echo.Context) error {
	var formData v1.ApplicationsDto
	id := context.QueryParam("companyId")
	if id == "" {
		return context.JSON(404, common.ResponseDTO{
			Message: "company id is required",
		})
	}
	repoId := context.QueryParam("repositoryId")
	if repoId == "" {
		return context.JSON(404, common.ResponseDTO{
			Message: "repository id is required",
		})
	}
	updateOption := context.QueryParam("companyUpdateOption")
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	var payload []v1.Application
	payload = formData.Applications
	for i, _ := range payload {
		payload[i].Url = UrlFormatter(payload[i].Url)
		if payload[i].MetaData.Labels == nil {
			payload[i].MetaData.Labels = make(map[string]string)
		}
		payload[i].MetaData.Labels["CompanyId"] = id
	}
	var options v1.ApplicationUpdateOption
	options.Option = enums.APPLICATION_UPDATE_OPTION(updateOption)
	err := a.companyService.UpdateApplications(id, repoId, payload, options)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	return common.GenerateSuccessResponse(context, payload,
		nil, "saved Successfully")
}

// NewApplicationApi returns Application type api
func NewApplicationApi(companyService service.Company, observerList []service.Observer) api.Application {
	return &applicationApi{
		companyService: companyService,
		observerList:   observerList,
	}
}
