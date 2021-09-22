package in_memory

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/enums"
)

var (
	CompanyCollection = "CompanyCollection"
)

type companyRepository struct {
}

func (c companyRepository) AppendRepositories(companyId string, repos []v1.Repository) error {
	var companies []v1.Company
	for _, each := range IndexedCompanies {
		companies = append(companies, each)
	}
	for _, each := range companies {
		if companyId == each.Id {
			for _, eachRepo := range repos {
				each.Repositories = append(each.Repositories, eachRepo)
			}
		}
	}
	return nil
}

func (c companyRepository) DeleteRepositories(companyId string, repos []v1.Repository, isSoftDelete bool) error {
	var companies []v1.Company
	var repositories []v1.Repository
	for _, each := range IndexedCompanies {
		companies = append(companies, each)
	}
	for _, each := range companies {
		if companyId == each.Id {
			if isSoftDelete {
				each.Status = enums.INACTIVE
			} else {
				for i, eachRepo := range each.Repositories {
					for _, DeleteRepo := range repos {
						if eachRepo.Id == DeleteRepo.Id {
							repositories = RemoveRepository(each.Repositories, i)
						}
					}
				}
				each.Repositories = repositories
			}
		}
	}
	return nil
}

func (c companyRepository) AppendApplications(companyId, repositoryId string, apps []v1.Application) error {
	var companies []v1.Company
	for _, each := range IndexedCompanies {
		companies = append(companies, each)
	}
	for _, each := range companies {
		for _, eachRepo := range each.Repositories {
			if eachRepo.Id == repositoryId {
				for _, eachApp := range apps {
					eachRepo.Applications = append(eachRepo.Applications, eachApp)
				}
			}
		}
	}
	return nil
}

func (c companyRepository) DeleteApplications(companyId, repositoryId string, apps []v1.Application, isSoftDelete bool) error {
	var companies []v1.Company
	var applications []v1.Application
	for _, each := range IndexedCompanies {
		companies = append(companies, each)
	}
	for _, each := range companies {
		if companyId == each.Id {
			if isSoftDelete {
				each.Status = enums.INACTIVE
			} else {
				for _, eachRepo := range each.Repositories {
					if eachRepo.Id == repositoryId {
						for i, eachApp := range eachRepo.Applications {
							for _, deleteApp := range apps {
								if eachApp.MetaData.Id == deleteApp.MetaData.Id {
									applications = RemoveApplication(eachRepo.Applications, i)
								}
							}
						}
					}
					eachRepo.Applications = applications
				}
			}
		}
	}
	return nil
}

func (c companyRepository) GetRepositoryByCompanyIdAndApplicationUrl(id, url string) v1.Repository {
	var company v1.Company
	var result v1.Repository
	for _, each := range IndexedCompanies {
		if id == each.Id {
			company = each
		}
	}
	for _, eachRepo := range company.Repositories {
		for _, app := range eachRepo.Applications {
			if app.Url == url {
				result = eachRepo
			}
		}
	}
	return result
}

func (c companyRepository) GetCompanyByApplicationUrl(url string) v1.Company {
	var companies []v1.Company
	var result v1.Company
	for _, each := range IndexedCompanies {
		companies = append(companies, each)
	}
	for _, each := range companies {
		for _, eachRepo := range each.Repositories {
			for _, app := range eachRepo.Applications {
				if app.Url == url {
					result = each
				}
			}
		}
	}
	return result
}

func (c companyRepository) GetCompanies(option v1.CompanyQueryOption) ([]v1.Company, int64) {
	var companies []v1.Company
	var result []v1.Company
	for _, each := range IndexedCompanies {
		companies = append(companies, each)
	}
	for i, _ := range companies {
		if option.LoadRepositories {
			if option.LoadApplications {
				result = companies
			} else {
				for j, _ := range companies[i].Repositories {
					companies[i].Repositories[j].Applications = nil
				}
				result = append(result, companies[i])
			}
		} else {
			companies[i].Repositories = nil
			result = companies
		}
	}
	return result, int64(len(result))
}

func (c companyRepository) GetByCompanyId(id string, option v1.CompanyQueryOption) v1.Company {
	var companies v1.Company
	for _, each := range IndexedCompanies {
		if each.Id == id {
			companies = each
		}
	}
	if option.LoadRepositories {
		if option.LoadApplications {
			return companies
		} else {
			for j, _ := range companies.Repositories {
				companies.Repositories[j].Applications = nil
			}
		}
	} else {
		companies.Repositories = nil
	}
	return companies
}

func (c companyRepository) GetRepositoriesByCompanyId(id string, option v1.CompanyQueryOption) []v1.Repository {
	var repository []v1.Repository
	var companies v1.Company
	for _, each := range IndexedCompanies {
		if each.Id == id {
			companies = each
		}
	}
	if option.LoadRepositories {
		if option.LoadApplications {
			for j, _ := range companies.Repositories {
				repository = append(repository, companies.Repositories[j])
			}
		} else {
			for j, _ := range companies.Repositories {
				companies.Repositories[j].Applications = nil
			}
			repository = companies.Repositories
		}
	} else {
		companies.Repositories = nil
	}
	return repository
}

func (c companyRepository) GetApplicationsByCompanyId(id string, option v1.CompanyQueryOption) []v1.Application {
	var applications []v1.Application
	var companies v1.Company
	for _, each := range IndexedCompanies {
		if each.Id == id {
			companies = each
		}
	}
	if option.LoadRepositories {
		for j, _ := range companies.Repositories {
			if option.LoadApplications {
				applications = append(applications, companies.Repositories[j].Applications...)
			} else {
				for j, _ := range companies.Repositories {
					companies.Repositories[j].Applications = nil
				}
				applications = companies.Repositories[j].Applications
			}
		}

	} else {
		companies.Repositories = nil
	}
	return applications
}
func (c companyRepository) GetApplicationsByCompanyIdAndRepositoryType(id string, _type enums.REPOSITORY_TYPE, option v1.CompanyQueryOption) []v1.Application {
	var applications []v1.Application
	var companies v1.Company
	for _, each := range IndexedCompanies {
		if each.Id == id {
			companies = each
		}
	}
	if option.LoadRepositories {
		for j, _ := range companies.Repositories {
			if _type == companies.Repositories[j].Type {
				if option.LoadApplications {
					applications = append(applications, companies.Repositories[j].Applications...)
				} else {
					for j, _ := range companies.Repositories {
						companies.Repositories[j].Applications = nil
					}
					applications = companies.Repositories[j].Applications
				}
			}
		}

	} else {
		companies.Repositories = nil
	}
	return applications
}

func (c companyRepository) Store(company v1.Company) error {
	if len(IndexedCompanies) == 0 {
		IndexedCompanies = make(map[string]v1.Company)
	}
	IndexedCompanies[company.Id] = company
	return nil
}
func (c companyRepository) Update(company v1.Company, companyUpdateOption v1.CompanyUpdateOption) {
	panic("implement me")
}

func (c companyRepository) Delete(companyId string) error {
	delete(IndexedCompanies, companyId)
	return nil
}
func paginate(logs []v1.Company, page int64, limit int64) []v1.Company {
	if page < 0 || limit <= 0 {
		return nil
	}
	var startIndex, endIndex int64
	if page == 0 {
		startIndex = 0
	} else {
		startIndex = page * limit
	}
	endIndex = startIndex + limit
	if startIndex >= int64(len(logs)) {
		return nil
	}
	if endIndex >= int64(len(logs)) {
		return logs[startIndex:]
	}
	return logs[startIndex:endIndex]
}
func RemoveRepository(s []v1.Repository, i int) []v1.Repository {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
func RemoveApplication(s []v1.Application, i int) []v1.Application {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func NewCompanyRepository(timeout int) repository.CompanyRepository {
	return &companyRepository{}

}
