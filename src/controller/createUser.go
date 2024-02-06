package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/rest_err"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestErr("Você chamou a rota errada")
	c.JSON(err.Code, err)
}
