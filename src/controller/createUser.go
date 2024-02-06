package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/controller/model/request"
	"github.com/mathews-r/golang/src/rest_err"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestErr(fmt.Sprintf("Some fields are incorrect"))
		c.JSON(restErr.Code, restErr)
		return
	}
	c.JSON(200, userRequest)
}
