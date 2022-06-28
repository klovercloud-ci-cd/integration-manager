package v1

import (
	"errors"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"reflect"
)

// Application contains application info
type Application struct {
	MetaData     ApplicationMetadata  `bson:"_metadata" json:"_metadata"`
	RepositoryId string               `bson:"repositoryId" json:"repositoryId"`
	CompanyId    string               `bson:"companyId" json:"companyId"`
	Url          string               `bson:"url" json:"url"`
	Webhook      GitWebhook           `bson:"webhook" json:"webhook"`
	Status       enums.COMPANY_STATUS `bson:"status" json:"status"`
}

// ApplicationDto contains application info
type ApplicationDto struct {
	MetaData ApplicationMetadata  `bson:"_metadata" json:"_metadata"`
	Url      string               `bson:"url" json:"url"`
	Webhook  GitWebhook           `bson:"webhook" json:"webhook"`
	Status   enums.COMPANY_STATUS `bson:"status" json:"status"`
}

// Validate validates application info
//func (application Application) Validate() error {
//	if application.Url == "" {
//		return errors.New("Application url is required!")
//	}
//	err := application.MetaData.Validate()
//	if err != nil {
//		return err
//	}
//	return nil
//}

// ApplicationMetadata contains application metadata info
type ApplicationMetadata struct {
	Labels           map[string]string `bson:"labels" json:"labels"`
	Id               string            `bson:"id" json:"id"`
	Name             string            `bson:"name" json:"name"`
	IsWebhookEnabled bool              `bson:"is_webhook_enabled" json:"is_webhook_enabled"`
}

// ApplicationMetadataCollection contains application metadata collection info
type ApplicationMetadataCollection struct {
	MetaData ApplicationMetadata  `bson:"_metadata" json:"_metadata"`
	Status   enums.COMPANY_STATUS `bson:"status" json:"status"`
}

// Validate validates application metadata
func (metadata ApplicationMetadata) Validate() error {
	keys := reflect.ValueOf(metadata.Labels).MapKeys()
	for i := 0; i < len(keys); i++ {
		if metadata.Labels[keys[i].String()] == "" {
			return errors.New("Application metadata label is missing!")
		}
	}
	if metadata.Id == "" {
		return errors.New("Application metadata id is required!")
	}
	if metadata.Name == "" {
		return errors.New("Application metadata name is required!")
	}
	return nil
}

// ApplicationUpdateOption contains applications update options
type ApplicationUpdateOption struct {
	Option enums.APPLICATION_UPDATE_OPTION `json:"option"`
}

// ApplicationsDto contains application list
//type ApplicationsDto struct {
//	Applications []Application `bson:"applications" json:"applications"`
//}
