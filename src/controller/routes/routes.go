package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.GetUserById)
	r.GET("/getUserByEmail/:userEmail", controller.GetUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
