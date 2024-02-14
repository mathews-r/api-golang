package repository

import (
	"context"
	"os"

	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"github.com/mathews-r/golang/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	DB_CATEGORY_COLLECTION = "DB_CATEGORY_COLLECTION"
)

func (cr *categoryRepository) CreateCategory(categoryDomain model.CategoryDomainInterface) (model.CategoryDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createCategory repository")

	collectionName := os.Getenv(DB_CATEGORY_COLLECTION)

	collection := cr.databaseConnection.Collection(collectionName)

	value := converter.ConvertDomainToEntityCategory(categoryDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerErr(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)
	return converter.ConvertEntityToDomainCategory(*value), nil
}
