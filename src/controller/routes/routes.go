package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/controller"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", userController.GetUserById)
	r.GET("/getUserByEmail/:userEmail", userController.GetUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
}
