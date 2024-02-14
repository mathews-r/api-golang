package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/model/service"
)

func NewCategoryControllerInterface(serviceInterface service.CategoryDomainService) CategoryControllerInterface {
	return &categoryControllerInterface{serviceInterface}
}

type CategoryControllerInterface interface {
	CreateCategory(c *gin.Context)
}

type categoryControllerInterface struct {
	service service.CategoryDomainService
}
