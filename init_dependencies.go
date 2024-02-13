package main

import (
	"github.com/mathews-r/golang/src/controller"
	"github.com/mathews-r/golang/src/model/repository"
	"github.com/mathews-r/golang/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)

	return controller.NewUserControllerInterface(service)
}
