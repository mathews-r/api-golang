package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mathews-r/golang/src/configs/database/mongodb"
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/controller"
	"github.com/mathews-r/golang/src/controller/model/repository"
	"github.com/mathews-r/golang/src/controller/routes"
	"github.com/mathews-r/golang/src/model/service"
)

func main() {
	logger.Info("About to start user application")
	godotenv.Load()

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserDomainService(userRepo)
	userController := controller.NewUserControllerInterface(userService)

	postRepo := repository.NewPostRepository(database)
	postService := service.NewPostDomainService(postRepo)
	postController := controller.NewPostControllerInterface(postService)

	// userController := initDependencies(database)
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController, postController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
