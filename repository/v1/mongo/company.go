package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/enums"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	CompanyCollection = "CompanyCollection"
)

type companyRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (c companyRepository) GetRepositoryByRepositoryId(id string) v1.Repository {
	var repo v1.Repository
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"repositories.id": id}}
	query["$and"] = and
	coll := c.manager.Db.Collection(CompanyCollection)
	result, err := coll.Find(c.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValues := new(v1.Company)
		err := result.Decode(elemValues)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		for _, each := range elemValues.Repositories {
			repo = each
		}
	}
	return repo
}

func (c companyRepository) GetApplicationByCompanyIdAndRepositoryIdAndApplicationUrl(companyId, repositoryId, applicationUrl string) v1.Application {
	var app v1.Application
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"id": companyId, "repositories.id": repositoryId, "repositories.applications.url": applicationUrl}}
	query["$and"] = and
	coll := c.manager.Db.Collection(CompanyCollection)
	result, err := coll.Find(c.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValues := new(v1.Company)
		err := result.Decode(elemValues)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		for _, each := range elemValues.Repositories {
			for _, eachApp := range each.Applications {
				app = eachApp
			}
		}
	}
	return app
}

func (c companyRepository) AppendRepositories(companyId string, repos []v1.Repository) error {
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"id": companyId}}
	query["$and"] = and
	coll := c.manager.Db.Collection(CompanyCollection)
	result, err := coll.Find(c.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValues := new(v1.Company)
		err := result.Decode(elemValues)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		for _, eachRepo := range repos {
			elemValues.Repositories = append(elemValues.Repositories, eachRepo)
		}
		err2 := c.Store(*elemValues)
		if err2 != nil {
			return err2
		}
		break
	}
	return nil
}

func (c companyRepository) DeleteRepositories(companyId string, repos []v1.Repository, isSoftDelete bool) error {
	var repositories []v1.Repository
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"id": companyId}}
	query["$and"] = and
	coll := c.manager.Db.Collection(CompanyCollection)
	result, err := coll.Find(c.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValues := new(v1.Company)
		err := result.Decode(elemValues)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		if isSoftDelete {
			elemValues.Status = enums.INACTIVE
			err2 := c.Store(*elemValues)
			if err2 == nil {
				return nil
			}
		} else {
			for i, eachRepo := range elemValues.Repositories {
				for _, deleteRepo := range repos {
					if eachRepo.Id == deleteRepo.Id {
						repositories = RemoveRepository(elemValues.Repositories, i)
					}
				}
			}
			elemValues.Repositories = repositories
			er := c.Store(*elemValues)
			if er != nil {
				return er
			}
		}
	}
	return nil
}

func (c companyRepository) AppendApplications(companyId, repositoryId string, apps []v1.Application) error {
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"id": companyId, "repositories.id": repositoryId}}
	query["$and"] = and
	coll := c.manager.Db.Collection(CompanyCollection)
	result, err := coll.Find(c.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValues := new(v1.Company)
		err := result.Decode(elemValues)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		var repo v1.Repository
		for _, each := range elemValues.Repositories {
			repo = each
			repo.Applications = each.Applications
		}
		for _, each := range apps {
			repo.Applications = append(repo.Applications, each)
		}
		var com v1.Company
		com.Id = elemValues.Id
		com.MetaData = elemValues.MetaData
		com.Name = elemValues.Name
		com.Repositories = append(com.Repositories, repo)
		com.Status = elemValues.Status
		er := c.Store(com)
		if er != nil {
			return er
		}
	}
	return nil
}

func (c companyRepository) DeleteApplications(companyId, repositoryId string, apps []v1.Application, isSoftDelete bool) error {
	var applications []v1.Application
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"id": companyId}}
	query["$and"] = and
	coll := c.manager.Db.Collection(CompanyCollection)
	result, err := coll.Find(c.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValues := new(v1.Company)
		err := result.Decode(elemValues)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		if isSoftDelete {
			elemValues.Status = enums.INACTIVE
			err2 := c.Store(*elemValues)
			if err2 == nil {
				return nil
			}
		} else {
			for _, eachRepo := range elemValues.Repositories {
				if repositoryId == eachRepo.Id {
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
			err2 := c.Store(*elemValues)
			if err2 != nil {
				return err2
			}
			break
		}
	}
	return nil
}

func (c companyRepository) GetRepositoryByCompanyIdAndApplicationUrl(id, url string) v1.Repository {
	var results v1.Repository
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"id": id}}
	query["$and"] = and
	coll := c.manager.Db.Collection(CompanyCollection)
	result, err := coll.Find(c.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Company)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		for _, each := range elemValue.Repositories {
			for _, eachApp := range each.Applications {
				if url == eachApp.Url {
					results = each
				}
			}
		}
	}
	return results
}

func (c companyRepository) GetCompanyByApplicationUrl(url string) v1.Company {
	var results v1.Company
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"repositories.applications.url": url}}
	query["$and"] = and
	coll := c.manager.Db.Collection(CompanyCollection)
	result, err := coll.Find(c.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Company)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = *elemValue
	}
	return results
}

func (c companyRepository) GetCompanies(option v1.CompanyQueryOption) ([]v1.Company, int64) {
	var results []v1.Company
	coll := c.manager.Db.Collection(CompanyCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	result, err := coll.Find(c.manager.Ctx, bson.D{}, &options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
	})
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Company)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		if option.LoadRepositories {
			if option.LoadApplications {
				results = append(results, *elemValue)
			} else {
				results = append(results, elemValue.GetCompanyWithRepository())
			}
		} else {
			results = append(results, elemValue.GetCompanyWithoutRepository())
		}
	}
	return results, int64(len(results))
}

func (c companyRepository) GetByCompanyId(id string, option v1.CompanyQueryOption) (v1.Company, int64) {
	var results v1.Company
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"id": id}}
	query["$and"] = and
	coll := c.manager.Db.Collection(CompanyCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	result, err := coll.Find(c.manager.Ctx, query, &options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
	})
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Company)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		if option.LoadRepositories {
			if option.LoadApplications {
				results = *elemValue
			} else {
				results = elemValue.GetCompanyWithRepository()
			}
		} else {
			results = elemValue.GetCompanyWithoutRepository()
		}
	}
	count, err := coll.CountDocuments(c.manager.Ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	return results, count
}

func (c companyRepository) GetRepositoriesByCompanyId(id string, option v1.CompanyQueryOption) ([]v1.Repository, int64) {
	var results []v1.Repository
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"id": id}}
	query["$and"] = and
	coll := c.manager.Db.Collection(CompanyCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	result, err := coll.Find(c.manager.Ctx, query, &options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
	})
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Company)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		if option.LoadRepositories == true {
			if option.LoadApplications == true {
				results = elemValue.Repositories
			} else {
				var rep v1.Repository
				for _, each := range elemValue.Repositories {
					rep.Type = each.Type
					rep.Id = each.Id
					rep.Token = each.Token
					rep.Applications = nil

					results = append(results, rep)
				}
			}
			return results, int64(len(results))
		}
	}
	return results, int64(len(results))
}

func (c companyRepository) GetApplicationsByCompanyId(id string, option v1.CompanyQueryOption) ([]v1.Application, int64) {
	var company v1.Company
	var results []v1.Application
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"id": id}}
	query["$and"] = and
	coll := c.manager.Db.Collection(CompanyCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	result, err := coll.Find(c.manager.Ctx, query, &options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
	})
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Company)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		for _, eachRepo := range elemValue.Repositories {
			for _, eachApp := range eachRepo.Applications {
				results = append(results, eachApp)
			}
		}
		company = *elemValue
	}
	return results, int64(len(company.Repositories))
}

func (c companyRepository) GetApplicationsByCompanyIdAndRepositoryType(id string, _type enums.REPOSITORY_TYPE, option v1.CompanyQueryOption) []v1.Application {
	var results []v1.Application
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"id": id}}
	query["$and"] = and
	coll := c.manager.Db.Collection(CompanyCollection)
	skip := option.Pagination.Page * option.Pagination.Limit
	result, err := coll.Find(c.manager.Ctx, query, &options.FindOptions{
		Limit: &option.Pagination.Limit,
		Skip:  &skip,
	})
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Company)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		var app []v1.Application
		for _, each := range elemValue.Repositories {
			if _type == each.Type {
				app = each.Applications
			}
		}
		results = app
	}
	return results
}

func (c companyRepository) Store(company v1.Company) error {
	coll := c.manager.Db.Collection(CompanyCollection)
	_, err := coll.InsertOne(c.manager.Ctx, company)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
	return nil
}

func (c companyRepository) Delete(companyId string) error {
	coll := c.manager.Db.Collection(CompanyCollection)
	filter := bson.M{"id": companyId, "status": bson.M{"$in": enums.ACTIVE}}

	update := bson.M{"$set": bson.M{"status": enums.INACTIVE}}

	_, err := coll.UpdateOne(
		c.manager.Ctx,
		filter,
		update,
	)

	return err
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
	return &companyRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}

}
