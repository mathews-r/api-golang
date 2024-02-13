package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/model/service"
)

func NewPostControllerInterface(serviceInterface service.PostDomainService) PostControllerInterface {
	return &postControllerInterface{service: serviceInterface}
}

type PostControllerInterface interface {
	CreatePost(c *gin.Context)
	GetPosts(c *gin.Context)
	// DeletePost(c *gin.Context)
	// UpdatePost(c *gin.Context)
	// GetPostById(c *gin.Context)
}

type postControllerInterface struct {
	service service.PostDomainService
}
