package v1

import (
	"errors"
	"github.com/klovercloud-ci-cd/integration-manager/api/common"
	"github.com/klovercloud-ci-cd/integration-manager/config"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/api"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"github.com/labstack/echo/v4"
	"github.com/twinj/uuid"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"log"
	"strings"
)

type v1GithubApi struct {
	gitService                   service.Git
	companyService               service.Company
	processInventoryEventService service.ProcessInventoryEvent
	observerList                 []service.Observer
}

// GetBranches... Get Branches
// @Summary Get Branches
// @Description Gets Branches
// @Tags Github
// @Produce json
// @Param userName query string true "User Name"
// @Param repoId query string true "Repository Id"
// @Param companyId query string true "company Id"
// @Param repoName query string true "Repository Name"
// @Param loadApplications query bool false "Loads ApplicationsDto"
// @Param loadToken query bool true "Loads Token"
// @Success 200 {object} common.ResponseDTO{data=[]v1.GitBranches}
// @Router /api/v1/githubs [GET]
func (g v1GithubApi) GetBranches(context echo.Context) error {
	repoId := context.QueryParam("repoId")
	option := v1.CompanyQueryOption{
		Pagination:       v1.Pagination{},
		LoadRepositories: true,
		LoadApplications: true,
		LoadToken:        true,
	}
	id := context.QueryParam("companyId")
	repo := g.companyService.GetRepositoryByRepositoryId(id, repoId, option)
	userName := context.QueryParam("userName")
	if userName == "" {
		return errors.New("userName is required")
	}
	repoName := context.QueryParam("repoName")
	if repoName == "" {
		return errors.New("repoName is required")
	}
	branches, err := g.gitService.GetBranches(userName, repoName, repo.Token)
	if err != nil {
		return err
	}
	return common.GenerateSuccessResponse(context, branches, nil, "success")
}

// Listen ... Listen Github Web hook event
// @Summary  Listen Github Web hook event
// @Description Listens Github Web hook events. Register this endpoint as github web hook endpoint
// @Tags Github
// @Accept json
// @Produce json
// @Param data body v1.GithubWebHookEvent true "GithubWebHookEvent Data"
// @Success 200 {object} common.ResponseDTO{data=string}
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/githubs [POST]
func (g v1GithubApi) ListenEvent(context echo.Context) error {
	resource := new(v1.GithubWebHookEvent)
	if err := context.Bind(resource); err != nil {
		log.Println(err.Error())
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	companyId := context.QueryParam("companyId")
	if companyId == "" {
		return common.GenerateErrorResponse(context, "[ERROR] no companyId is provided", "Please provide companyId")
	}
	repoName := resource.Repository.Name
	owner := resource.Repository.Owner.Login
	revision := resource.After
	repository := g.companyService.GetRepositoryByCompanyIdAndApplicationUrl(companyId, resource.Repository.URL)
	application := g.companyService.GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(companyId, repository.Id, resource.Repository.URL)
	if !application.MetaData.IsWebhookEnabled {
		return common.GenerateForbiddenResponse(context, "[Forbidden]: Web hook is disabled!", "Operation Failed!")
	}
	data, err := g.gitService.GetPipeline(repoName, owner, revision, repository.Token)
	if err != nil {
		log.Println("[ERROR]:Failed to trigger pipeline process! ", err.Error())
		return common.GenerateErrorResponse(context, err.Error(), "Failed to trigger pipeline process!")
	}
	checkingFlag := BranchExists(data.Steps, resource.Ref, "GITHUB")
	if !checkingFlag {
		return common.GenerateErrorResponse(context, "Branch does not exist!", "Operation Failed!")
	}
	if data != nil {

		stepsCount := len(data.Steps)

		for i := 0; i < stepsCount; i++ {
			if data.Steps[i].Type == enums.BUILD {
				if images, ok := data.Steps[i].Params[enums.IMAGE]; ok {
					data.Steps[i].Params[enums.IMAGE] = setImageVersionForBuild(data.Steps[i], revision, images)
				}
				if storage, ok := data.Steps[i].Params[enums.STORAGE]; ok {
					data.Steps[i].Params[enums.STORAGE] = storage
				} else {
					data.Steps[i].Params[enums.STORAGE] = "500Mi"
				}
				if accessMode, ok := data.Steps[i].Params[enums.ACCESS]; ok {
					data.Steps[i].Params[enums.ACCESS] = setAccessModeForBuild(accessMode)
				}
				if buildType, ok := data.Steps[i].Params[enums.BUILD_TYPE]; ok {
					data.Steps[i].Params[enums.BUILD_TYPE] = buildType
				}

			} else if data.Steps[i].Type == enums.DEPLOY {

				isThisStepValidForThisCommit := false
				if data.Steps[i].Params[enums.REVISION] != "" {
					allowedRevisions := strings.Split(data.Steps[i].Params[enums.REVISION], ",")
					branch := strings.Split(resource.Ref, "/")[2]
					for _, each := range allowedRevisions {
						if each == branch {
							isThisStepValidForThisCommit = true
							break
						}
					}
				}
				if isThisStepValidForThisCommit {
					data.Steps[i].Params["images"] = setDeploymentVersion(data.Steps[i], revision)
					descriptor := g.setDescriptors(data.Steps[i], repoName, owner, revision, repository.Token)
					if descriptor != nil {
						data.Steps[i].Descriptors = descriptor
					} else {
						return common.GenerateErrorResponse(context, err.Error(), "Failed to trigger pipeline process!")
					}
				} else {
					data.Steps = append(data.Steps[:i], data.Steps[i+1:]...)
					stepsCount = stepsCount - 1
					i = i - 1
				}
			} else if data.Steps[i].Type == enums.INTERMEDIARY {
				if images, ok := data.Steps[i].Params["images"]; ok {
					data.Steps[i].Params["images"] = setImageVersionForIntermediary(data.Steps[i], revision, images)
				}
			}
		}
	}
	data.ProcessId = uuid.NewV4().String()

	company, _ := g.companyService.GetByCompanyId(companyId, v1.CompanyQueryOption{v1.Pagination{}, false, false, false})
	todaysRanProcess := g.processInventoryEventService.CountTodaysRanProcessByCompanyId(companyId)
	data.MetaData = v1.PipelineMetadata{
		CompanyId:       companyId,
		CompanyMetadata: company.MetaData,
		CommitId:        revision,
	}
	subject := v1.Subject{
		Log:                   "Pipeline triggered",
		CoreRequestQueryParam: map[string]string{"url": resource.Repository.URL, "revision": revision, "purging": config.PipelinePurging},
		EventData:             map[string]interface{}{},
		Pipeline:              *data,
		App: struct {
			CompanyId    string
			AppId        string
			RepositoryId string
		}{
			CompanyId:    companyId,
			AppId:        application.MetaData.Id,
			RepositoryId: repository.Id,
		},
	}
	if todaysRanProcess >= company.MetaData.TotalProcessPerDay {
		subject.Log = "No More process today, you've touched today's limit!"
		if subject.EventData == nil {
			subject.EventData = make(map[string]interface{})
		}
		subject.EventData["trigger"] = false
		subject.EventData["log"] = subject.Log
	}

	go g.notifyAll(subject)
	return common.GenerateSuccessResponse(context, data.ProcessId, nil, "Pipeline triggered!")
}

// setImageVersionForBuild returns image version for build step
func setImageVersionForBuild(step v1.Step, revision string, images string) string {
	imageRevision := revision
	if step.Params[enums.REVISION] != "" {
		imageRevision = step.Params[enums.REVISION]
	}
	listOfImages := strings.Split(images, ",")
	for i, image := range listOfImages {
		strs := strings.Split(image, ":")
		if len(strs) == 1 {
			listOfImages[i] = listOfImages[i] + ":" + imageRevision
		}
	}
	return strings.Join(listOfImages, ",")
}

// setAccessModeForBuild returns access mode for build step
func setAccessModeForBuild(accessMode string) string {
	if accessMode == string(enums.READ_WRITE_ONCE_POD) {
		return "ReadWriteOncePod"
	} else if accessMode == string(enums.READ_WRITE_MANY) {
		return "ReadWriteMany"
	} else if accessMode == string(enums.READ_ONLY_MANY) {
		return "ReadOnlyMany"
	} else {
		return "ReadWriteOnce"
	}
}

// setImageVersionForIntermediary returns image version for Intermediary step
func setImageVersionForIntermediary(step v1.Step, revision string, img string) string {
	images := strings.Split(img, ",")
	imageRevision := revision
	if step.Params[enums.REVISION] != "" {
		imageRevision = step.Params[enums.REVISION]
	}
	for i, image := range images {
		strs := strings.Split(image, ":")
		if len(strs) == 1 {
			images[i] = images[i] + ":" + imageRevision
		}
	}
	return strings.Join(images, ",")
}

// setDeploymentVersion returns image version for deployment
func setDeploymentVersion(step v1.Step, revision string) string {
	var deploymentVersion string
	if images, ok := step.Params["images"]; ok {
		images := strings.Split(images, ",")
		for i, image := range images {
			strs := strings.Split(image, ":")
			if len(strs) == 1 {
				if step.Params["trunk_based"]=="enabled"{
					images[i] = images[i] + ":" + revision
				}else{
					images[i] = images[i]
				}
			}
		}
		deploymentVersion = strings.Join(images, ",")
	}
	return deploymentVersion
}

// setDescriptors returns descriptors for deployment
func (g v1GithubApi) setDescriptors(step v1.Step, repoName string, owner string, revision string, token string) *[]unstructured.Unstructured {
	var descriptor *[]unstructured.Unstructured
	if val, ok := step.Params["env"]; ok {
		contentsData, err := g.gitService.GetDescriptors(repoName, owner, revision, token, enums.PIPELINE_DESCRIPTORS_BASE_DIRECTORY+"/", val)
		if err != nil {
			return nil
		}
		if contentsData != nil {
			descriptor = &contentsData
		}
	}
	return descriptor
}
func (g v1GithubApi) notifyAll(listener v1.Subject) {
	for _, observer := range g.observerList {
		go observer.Listen(listener)
	}
}

// NewGithubApi returns Git type api
func NewGithubApi(gitService service.Git, companyService service.Company, processInventoryEventService service.ProcessInventoryEvent, observerList []service.Observer) api.Git {
	return &v1GithubApi{
		gitService:                   gitService,
		companyService:               companyService,
		observerList:                 observerList,
		processInventoryEventService: processInventoryEventService,
	}
}
