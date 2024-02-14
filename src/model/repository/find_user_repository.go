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
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID repository",
		zap.String("journey", "findUserByID"))

	collectionName := os.Getenv(DB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findUserByID"))
			return nil, rest_err.NewNotFoundErr(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserByID"))
		return nil, rest_err.NewInternalServerErr(errorMessage)
	}
	logger.Info("FindUserByID repository executed successfully",
		zap.String("journey", "findUserByID"),
		zap.String("userId", userEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(email string, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmailAndPassword repository",
		zap.String("journey", "findUserByEmailAndPassword"))

	collectionName := os.Getenv(DB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}, {Key: "password", Value: password}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User or password is invalid: %s", email)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findUserByEmailAndPassword"))
			return nil, rest_err.NewNotFoundErr(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserByEmailAndPassword"))

		return nil, rest_err.NewInternalServerErr(errorMessage)
	}

	logger.Info("FindUserByEmailAndPassword repository executed successfully",
		zap.String("journey", "findUserByEmailAndPassword"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(*userEntity), nil
}
