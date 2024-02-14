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
	DB_POST_COLLECTION = "DB_POST_COLLECTION"
)

func (ur *postRepository) CreatePost(postDomain model.PostDomainInterface) (model.PostDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreatePost repository",
		zap.String("journey", "CreatePost"))

	collectionName := os.Getenv(DB_POST_COLLECTION)

	collection := ur.databaseConnection.Collection(collectionName)

	value := converter.ConvertDomainToEntityPost(postDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to create post",
			err,
			zap.String("journey", "CreatePost"))
		return nil, rest_err.NewInternalServerErr(err.Error())
	}

	logger.Info(
		"CreatePost repository executed successfully",
		zap.String("postId", value.ID.Hex()),
		zap.String("journey", "CreatePost"))

	value.ID = result.InsertedID.(primitive.ObjectID)
	return converter.ConvertEntityToDomainPost(*value), nil
}
