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

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init loginUser controller",
		zap.String("journey", "loginUser"),
	)

	var userLogin request.UserLogin

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "loginUser"))

		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.LoginDomain(
		userLogin.Email,
		userLogin.Password,
	)

	domainResult, token, err := uc.service.LoginUser(domain)
	if err != nil {
		logger.Error(
			"Error trying to call loginUser service",
			err,
			zap.String("journey", "loginUser"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"loginUser controller executed successfully",
		zap.String("userId", domainResult.GetId()),
		zap.String("journey", "loginUser"))

	c.Header("Authorization", token)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
