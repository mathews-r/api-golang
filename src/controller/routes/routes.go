package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/controller"
	"github.com/mathews-r/golang/src/model"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface, postController controller.PostControllerInterface, categoryController controller.CategoryControllerInterface) {

	r.GET("/getUserByEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.POST("/login", userController.LoginUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)

	// POST ROUTES
	r.POST("/posts/", model.VerifyTokenMiddleware, postController.CreatePost)
	r.GET("/posts/", model.VerifyTokenMiddleware, postController.GetPosts)
	r.GET("/posts/:postId", model.VerifyTokenMiddleware, postController.GetPostById)
	r.PUT("/posts/:postId", model.VerifyTokenMiddleware, postController.UpdatePost)
	r.DELETE("/posts/:postId", model.VerifyTokenMiddleware, postController.DeletePost)
	// r.GET("/posts/search", model.VerifyTokenMiddleware, postController.GetPostByQuery)

	// CATEGORY ROUTES
	r.POST("/categories/", model.VerifyTokenMiddleware, categoryController.CreateCategory)
	// r.GET("/categories/", model.VerifyTokenMiddleware, categoryController.GetCategories)
}
