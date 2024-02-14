package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init deleteUser controller",
		zap.String("journey", "deleteUser"),
	)

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestErr("Invalid userid, must be hex")
		c.JSON(errRest.Code, errRest)
	}

	err := uc.service.DeleteUser(userId)
	if err != nil {
		logger.Error(
			"Error trying to call deleteUser service",
			err,
			zap.String("journey", "deleteUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"deleteUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"))

	c.Status(http.StatusOK)
}
