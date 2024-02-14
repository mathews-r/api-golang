package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"github.com/mathews-r/golang/src/model/repository/entity"
	"github.com/mathews-r/golang/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (pr *postRepository) GetPosts() ([]model.PostDomainInterface, *rest_err.RestErr) {
	logger.Info("Init GetPosts repository")

	collectionName := os.Getenv(DB_POST_COLLECTION)

	collection := pr.databaseConnection.Collection(collectionName)

	// postEntity := &entity.PostEntity{}

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		errorMessage := "Error finding posts in database"
		logger.Error(errorMessage, err)
		return nil, rest_err.NewInternalServerErr(errorMessage)
	}
	defer cursor.Close(context.Background())

	var posts []model.PostDomainInterface
	for cursor.Next(context.Background()) {
		var postEntity entity.PostEntity
		if err := cursor.Decode(&postEntity); err != nil {
			logger.Error("Error decoding post", err)
			continue
		}
		posts = append(posts, converter.ConvertEntityToDomainPost(postEntity))
	}

	if err := cursor.Err(); err != nil {
		errorMessage := "Error iterating through posts"
		logger.Error(errorMessage, err)
		return nil, rest_err.NewInternalServerErr(errorMessage)
	}

	logger.Info("Executed successfully")
	return posts, nil
}

func (pr *postRepository) GetPostById(postId string) (model.PostDomainInterface, *rest_err.RestErr) {
	logger.Info("Init GetPostById repository",
		zap.String("journey", "GetPostById"))

	collectionName := os.Getenv(DB_POST_COLLECTION)

	collection := pr.databaseConnection.Collection(collectionName)

	postEntity := &entity.PostEntity{}

	objectId, _ := primitive.ObjectIDFromHex(postId)
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(context.Background(), filter).Decode(postEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("post not found in database: %s", postId)
			logger.Error(errorMessage, err)
			return nil, rest_err.NewNotFoundErr(errorMessage)
		}
		errorMessage := "Error trying to find a post by id"
		logger.Error(errorMessage, err)
		return nil, rest_err.NewInternalServerErr(errorMessage)
	}

	logger.Info("GetPostById repository executed successfully",
		zap.String("journey", "GetPostById"),
		zap.String("postId", postId))

	return converter.ConvertEntityToDomainPost(*postEntity), nil
}
