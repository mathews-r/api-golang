package repository

import (
	"context"
	"os"

	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"github.com/mathews-r/golang/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (pr *postRepository) UpdatePost(postId string, postDomain model.PostDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateUser repository")

	collectionName := os.Getenv(DB_POST_COLLECTION)

	collection := pr.databaseConnection.Collection(collectionName)

	value := converter.ConvertDomainToEntityPost(postDomain)
	postIdHex, _ := primitive.ObjectIDFromHex(postId)

	filter := bson.D{{Key: "_id", Value: postIdHex}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return rest_err.NewInternalServerErr(err.Error())
	}

	logger.Info("Post updated successfully", zap.String("postId", postId))
	return nil
}
