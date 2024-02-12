package repository

import (
	"context"
	"os"

	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"github.com/mathews-r/golang/src/model/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *postRepository) CreatePost(postDomain model.PostDomainInterface) (model.PostDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository")

	collectionName := os.Getenv(DB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collectionName)

	value := converter.ConvertDomainToEntityPost(postDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerErr(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)
	return converter.ConvertEntityToDomainPost(*value), nil
}
