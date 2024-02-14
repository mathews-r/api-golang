package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/validation"
	"github.com/mathews-r/golang/src/controller/model/request"
	"github.com/mathews-r/golang/src/model"
	"github.com/mathews-r/golang/src/view"
	"go.uber.org/zap"
)

var (
	PostDomainInterface model.PostDomainInterface
)

func (pc *postControllerInterface) CreatePost(c *gin.Context) {
	var postRequest request.PostRequest

	if err := c.ShouldBindJSON(&postRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
	}

	token := c.Request.Header.Get("Authorization")

	user, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	domain := model.NewPostDomain(
		postRequest.Title,
		postRequest.Content,
		postRequest.Category,
		user.GetId(),
		postRequest.Published,
		postRequest.Updated,
	)

	domainResult, err := pc.service.CreatePost(domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Post created successfully", zap.String("journey", "createPost"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponsePost(domainResult))
}
