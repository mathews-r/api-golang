package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/configs/logger"
)

func (uc *userControllerInterface) GetUserById(c *gin.Context) {
	logger.Info("Usuario")
}

func (uc *userControllerInterface) GetUserByEmail(c *gin.Context) {

}
