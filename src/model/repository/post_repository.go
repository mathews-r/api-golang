package repository

import (
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type postRepository struct {
	databaseConnection *mongo.Database
}

type PostRepository interface {
	CreatePost(postDomain model.PostDomainInterface) (model.PostDomainInterface, *rest_err.RestErr)
	GetPosts() ([]model.PostDomainInterface, *rest_err.RestErr)
	GetPostById(string) (model.PostDomainInterface, *rest_err.RestErr)
	UpdatePost(string, model.PostDomainInterface) *rest_err.RestErr
	DeletePost(string) *rest_err.RestErr
}

func NewPostRepository(database *mongo.Database) PostRepository {
	return &postRepository{
		database,
	}
}
