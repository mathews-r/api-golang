package controller

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"github.com/mathews-r/golang/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userControllerInterface) GetUserById(c *gin.Context) {
	logger.Info("id")

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := rest_err.NewBadRequestErr("Invalid userid")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	userDomain, err := uc.service.FindUserByEmail(userId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))

}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("email")

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("User authenticated: %#v", user))

	email := c.Param("userEmail")

	if _, err := mail.ParseAddress(email); err != nil {
		errorMessage := rest_err.NewBadRequestErr("Invalid email")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	userDomain, err := uc.service.FindUserByEmail(email)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
