package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (pc *postControllerInterface) DeletePost(c *gin.Context) {
	logger.Info("Init DeletePost controller",
		zap.String("journey", "DeletePost"),
	)

	postId := c.Param("postId")

	if _, err := primitive.ObjectIDFromHex(postId); err != nil {
		logger.Error(
			"Error trying to call DeletePost service",
			err,
			zap.String("journey", "DeletePost"))

		errRest := rest_err.NewBadRequestErr("Invalid postid, must be hex")
		c.JSON(errRest.Code, errRest)
	}

	err := pc.service.DeletePost(postId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"DeletePost controller executed successfully",
		zap.String("postId", postId),
		zap.String("journey", "DeletePost"))

	c.Status(http.StatusOK)
}
