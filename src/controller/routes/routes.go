package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/controller"
	"github.com/mathews-r/golang/src/model"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface, postController controller.PostControllerInterface) {

	r.GET("/getUserByEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", model.VerifyTokenMiddleware, userController.CreateUser)
	r.POST("/login", userController.LoginUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)

	// POST ROUTES
	// r.GET("/posts/search", tokenValidation, postController.GetPostByQuery)
	r.GET("/posts/", postController.GetPosts)
	// r.GET("/posts/:postId", tokenValidation, postController)
	r.POST("/posts/", postController.CreatePost)
	// r.PUT("/posts/:postId", tokenValidation, postController.UpdatePost)
	// r.DELETE("/posts/:postId", tokenValidation, postController.DeletePost)

	// CATEGORY ROUTES
	// r.GET("/categories/", tokenValidation, categoryController.GetCategories)
	// r.POST("/categories/", tokenValidation, categoryController.NewCategory)

	// AUTHY ROUTES
	// R.POST("/auth", authController.Verify)
}
