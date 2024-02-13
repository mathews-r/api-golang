package repository

import (
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DB_USER_COLLECTION = "DB_USER_COLLECTION"
)

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(userId string) *rest_err.RestErr
}

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		database,
	}
}
