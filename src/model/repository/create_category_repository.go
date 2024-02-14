package repository

import (
	"context"
	"os"

	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"github.com/mathews-r/golang/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

var (
	DB_CATEGORY_COLLECTION = "DB_CATEGORY_COLLECTION"
)

func (cr *categoryRepository) CreateCategory(categoryDomain model.CategoryDomainInterface) (model.CategoryDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateCategory repository",
		zap.String("journey", "CreateCategory"))

	collectionName := os.Getenv(DB_CATEGORY_COLLECTION)

	collection := cr.databaseConnection.Collection(collectionName)

	value := converter.ConvertDomainToEntityCategory(categoryDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to create category",
			err,
			zap.String("journey", "CreateCategory"))
		return nil, rest_err.NewInternalServerErr(err.Error())
	}

	logger.Info(
		"CreateCategory repository executed successfully",
		zap.String("categoryId", value.ID.Hex()),
		zap.String("journey", "CreateCategory"))

	value.ID = result.InsertedID.(primitive.ObjectID)
	return converter.ConvertEntityToDomainCategory(*value), nil
}
