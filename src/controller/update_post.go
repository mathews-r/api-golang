package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/configs/validation"
	"github.com/mathews-r/golang/src/controller/model/request"
	"github.com/mathews-r/golang/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (pc *postControllerInterface) UpdatePost(c *gin.Context) {
	logger.Info("Init UpdatePost controller",
		zap.String("journey", "UpdatePost"))

	var postRequest request.PostUpdateRequest

	if err := c.ShouldBindJSON(&postRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "UpdatePost"))

		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	postId := c.Param("postId")
	if _, err := primitive.ObjectIDFromHex(postId); err != nil {
		errRest := rest_err.NewBadRequestErr("Invalid postId, must be hex")
		c.JSON(errRest.Code, errRest)
	}

	domain := model.UpdatePostDomain(
		postRequest.Title,
		postRequest.Content,
		postRequest.Category,
	)

	err := pc.service.UpdatePost(postId, domain)
	if err != nil {
		logger.Error(
			"Error trying to call UpdatePost service",
			err,
			zap.String("journey", "UpdatePost"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"UpdatePost controller executed successfully",
		zap.String("postId", postId),
		zap.String("journey", "UpdatePost"))

	c.Status(http.StatusOK)
}
