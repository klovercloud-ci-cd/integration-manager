package logic

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/integration-manager/api/common"
	"github.com/klovercloud-ci-cd/integration-manager/config"
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"log"
)

type processInventoryEventService struct {
	httpClient service.HttpClient
}

func (p processInventoryEventService) CountTodaysRanProcessByCompanyId(companyId string) int64 {
	url := config.EventStoreUrl + "/processes?companyId=" + companyId + "&operation=countTodaysProcessByCompanyId"
	header := make(map[string]string)
	header["token"] = config.Token
	header["Content-Type"] = "application/json"
	var count int64
	data, err := p.httpClient.Get(url, header)
	if err != nil {
		log.Println(err.Error())
		return count
	}
	response := common.ResponseDTO{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		log.Println(err.Error())
		return count
	}
	b, err := json.Marshal(response.Data)
	if err != nil {
		log.Println(err.Error())
		return count
	}
	err = json.Unmarshal(b, &count)
	if err != nil {
		log.Println(err.Error())
		return count
	}
	return count
}

func (p processInventoryEventService) Listen(subject v1.Subject) {
	if subject.App.CompanyId == "" {
		return
	}
	url := config.EventStoreUrl + "/processes"

	header := make(map[string]string)
	header["token"] = config.Token
	header["Content-Type"] = "application/json"

	process := v1.Process{
		ProcessId:    subject.Pipeline.ProcessId,
		CompanyId:    subject.App.CompanyId,
		AppId:        subject.App.AppId,
		RepositoryId: subject.App.RepositoryId,
		CommitId:     subject.Pipeline.MetaData.CommitId,
		Branch:       subject.App.Branch,
		Data:         nil,
	}
	b, err := json.Marshal(process)
	if err != nil {
		log.Println(err.Error())
		return
	}
	go p.httpClient.Post(url, header, b)
}

// NewProcessInventoryEventObserverService returns Observer type service
func NewProcessInventoryEventObserverService(httpPublisher service.HttpClient) service.Observer {
	return &processInventoryEventService{
		httpClient: httpPublisher,
	}
}

// NewProcessInventoryEventService returns ProcessInventoryEvent type service
func NewProcessInventoryEventService(httpPublisher service.HttpClient) service.ProcessInventoryEvent {
	return &processInventoryEventService{
		httpClient: httpPublisher,
	}
}
