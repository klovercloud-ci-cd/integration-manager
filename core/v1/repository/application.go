package repository

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
)

// ApplicationRepository company application repository related operations
type ApplicationRepository interface {
	GetAll(companyId string, option v1.CompanyQueryOption) ([]v1.Application, int64)
	GetById(companyId string, applicationId string) v1.Application
	GetByIdAndRepoId(companyId string, repoId string, applicationId string) v1.Application
	Store(application v1.Application) error
	StoreAll(applications []v1.Application) error
	GetByCompanyIdAndRepoId(companyId, repoId string, pagination bool, option v1.CompanyQueryOption, statusQuery bool, status v1.StatusQueryOption) ([]v1.Application, int64)
	GetApplicationsByCompanyIdAndRepositoryType(companyId string, _type enums.REPOSITORY_TYPE, pagination bool, option v1.CompanyQueryOption, statusQuery bool, status v1.StatusQueryOption) ([]v1.Application, int64)
	GetByCompanyIdAndRepositoryIdAndUrl(companyId, repositoryId, applicationUrl string) v1.Application
	GetByCompanyIdAndUrl(companyId, applicationUrl string) v1.Application
	Update(companyId, repositoryId string, application v1.Application) error
	SoftDeleteApplication(application v1.Application) error
	DeleteApplication(companyId, repositoryId, applicationId string) error
	SearchAppsByCompanyIdAndName(companyId, name string) []v1.Application
}
