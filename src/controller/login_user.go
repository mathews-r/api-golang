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
	var userLogin request.UserLogin

	if err := c.ShouldBindJSON(&userLogin); err != nil {
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
		c.JSON(err.Code, err)
		return
	}

	c.Header("Authorization", token)
	logger.Info("Login successfully", zap.String("journey", "LoginUser"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
