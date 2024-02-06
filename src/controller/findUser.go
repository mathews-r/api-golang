package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/configs/logger"
)

func GetUserById(c *gin.Context) {
	logger.Info("Usuario")
}

func GetUserByEmail(c *gin.Context) {

}
