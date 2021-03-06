package repository

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
)

// CompanyRepository company repository related operations
type CompanyRepository interface {
	Store(company v1.Company) error
	//AppendRepositories(companyId string, repos []v1.Repository) error
	//DeleteRepositories(companyId string, repos []v1.Repository) error
	//AppendApplications(companyId, repositoryId string, apps []v1.Application) error
	//DeleteApplications(companyId, repositoryId string, repos []v1.Repository) error
	Delete(companyId string) error
	GetCompanies(option v1.CompanyQueryOption, status v1.StatusQueryOption) ([]v1.Company, int64)
	GetByCompanyId(id string) v1.Company
	GetByName(name string, status v1.StatusQueryOption) v1.Company
	//GetRepositoriesByCompanyId(id string, option v1.CompanyQueryOption) ([]v1.Repository, int64)
	//GetApplicationsByRepositoryId(repoId string, companyId string, option v1.CompanyQueryOption, status v1.StatusQueryOption) ([]v1.Application, int64)
	//GetCompanyByApplicationUrl(url string) v1.Company
	//GetRepositoryByRepositoryId(id, repoId string, option v1.CompanyQueryOption) v1.Repository
	//GetApplicationsByCompanyIdAndRepositoryType(id string, _type enums.REPOSITORY_TYPE, option v1.CompanyQueryOption, status v1.StatusQueryOption) []v1.Application
	//GetRepositoryByCompanyIdAndApplicationUrl(id, url string) v1.Repository
	//GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(companyId, repositoryId, applicationUrl string) v1.Application
	//GetDashboardData(companyId string) v1.DashboardData
	//UpdateApplication(companyId string, repositoryId string, applicationId string, app v1.Application) error
	//GetApplicationByApplicationId(companyId string, repoId string, applicationId string) v1.Application
}
