package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/model/service"
)

func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{service: serviceInterface}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetUserByEmail(c *gin.Context)
	GetUserById(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
