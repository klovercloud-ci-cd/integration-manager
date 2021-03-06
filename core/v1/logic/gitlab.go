package logic

import (
	v1 "github.com/klovercloud-ci-cd/integration-manager/core/v1"
	"github.com/klovercloud-ci-cd/integration-manager/core/v1/service"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type gitlabService struct {
	observerList []service.Observer
	client       service.HttpClient
}

func (g gitlabService) GetCommitsByBranch(username, repositoryName, branch, token string, option v1.Pagination) ([]v1.GitCommit, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (g gitlabService) GetContent(repositoryName, username, token, path string) (v1.GitContent, error) {
	//TODO implement me
	panic("implement me")
}

func (g gitlabService) CreateDirectoryContent(repositoryName, username, token, path string, content v1.DirectoryContentCreatePayload) (v1.DirectoryContentCreateAndUpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g gitlabService) UpdateDirectoryContent(repositoryName, username, token, path string, content v1.DirectoryContentUpdatePayload) (v1.DirectoryContentCreateAndUpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (g gitlabService) GetCommitByBranch(username, repositoryName, branch, token string) (v1.GitCommit, error) {
	//TODO implement me
	panic("implement me")
}

func (g gitlabService) GetBranches(username, repositoryName, token string) (v1.GitBranches, error) {
	//TODO implement me
	panic("implement me")
}

func (g gitlabService) GetPipeline(repositoryName, username, revision, token string) (*v1.Pipeline, error) {
	panic("implement me")
}

func (g gitlabService) GetDescriptors(repositoryName, username, revision, token, path, env string) ([]unstructured.Unstructured, error) {
	panic("implement me")
}

func (g gitlabService) GetDirectoryContents(repositoryName, username, revision, token, path string) ([]v1.GitDirectoryContent, error) {
	panic("implement me")
}

func (g gitlabService) CreateRepositoryWebhook(username, repositoryName, token string, companyId string) (v1.GitWebhook, error) {
	panic("implement me")
}

func (g gitlabService) DeleteRepositoryWebhookById(username, repositoryName, webhookId, token string) error {
	panic("implement me")
}

// NewGitlabService returns Git type service
func NewGitlabService(observerList []service.Observer, client service.HttpClient) service.Git {
	return &gitlabService{
		observerList: observerList,
		client:       client,
	}
}
