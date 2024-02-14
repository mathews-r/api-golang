package service

import (
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"github.com/mathews-r/golang/src/model/repository"
)

func NewPostDomainService(postInterface repository.PostRepository) PostDomainService {
	return &postDomainService{postInterface}
}

type postDomainService struct {
	postRepository repository.PostRepository
}

type PostDomainService interface {
	CreatePost(model.PostDomainInterface) (model.PostDomainInterface, *rest_err.RestErr)
	GetPosts() ([]model.PostDomainInterface, *rest_err.RestErr)
	GetPostById(postId string) (model.PostDomainInterface, *rest_err.RestErr)
	UpdatePost(string, model.PostDomainInterface) *rest_err.RestErr
}
