package repository

import (
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type categoryRepository struct {
	databaseConnection *mongo.Database
}

type CategoryRepository interface {
	CreateCategory(categoryRepository model.CategoryDomainInterface) (model.CategoryDomainInterface, *rest_err.RestErr)
}

func NewCategoryRepository(database *mongo.Database) CategoryRepository {
	return &categoryRepository{
		database,
	}
}
