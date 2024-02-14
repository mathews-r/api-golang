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
	"go.uber.org/zap"
)

func (uc *userControllerInterface) GetUserById(c *gin.Context) {
	logger.Info("Init findUserByID controller",
		zap.String("journey", "findUserByID"),
	)

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userId",
			err,
			zap.String("journey", "findUserByID"),
		)
		errorMessage := rest_err.NewBadRequestErr(
			"UserID is not a valid id",
		)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmail(userId)
	if err != nil {
		logger.Error("Error trying to call findUserByID services",
			err,
			zap.String("journey", "findUserByID"),
		)

		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByID controller executed successfully",
		zap.String("journey", "findUserByID"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))

}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail controller",
		zap.String("journey", "findUserByEmail"),
	)

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))

	if err != nil {

		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("User authenticated: %#v", user))

	email := c.Param("userEmail")

	if _, err := mail.ParseAddress(email); err != nil {
		logger.Error("Error trying to validate userEmail",
			err,
			zap.String("journey", "findUserByEmail"),
		)
		errorMessage := rest_err.NewBadRequestErr("User email is not a valid email.")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmail(email)
	if err != nil {
		logger.Error("Error trying to call findUserByEmail services",
			err,
			zap.String("journey", "findUserByEmail"),
		)

		c.JSON(err.Code, err)
		return
	}

	logger.Info("findUserByEmail controller executed successfully",
		zap.String("journey", "findUserByEmail"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
