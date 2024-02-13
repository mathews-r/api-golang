package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/controller"
	"github.com/mathews-r/golang/src/model"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface, postController controller.PostControllerInterface) {

	//USER ROUTES
	// r.GET("/getUserById/:userId", userController.FindUserByEmailService)
	// r.GET("/users", tokenValidation, userController.GetUsers)
	// r.GET("/users/:userId", tokenValidation, userController.GetUserById)
	// r.DELETE("/users/:userId", tokenValidation, userController.DeleteUser)
	// r.POST("/users", userController.CreateUser)

	r.GET("/getUserByEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", model.VerifyTokenMiddleware, userController.CreateUser)
	r.POST("/login", userController.LoginUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)

	// POST ROUTES
	// r.GET("/posts/search", tokenValidation, postController.GetPostByQuery)
	// r.GET("/posts/", tokenValidation, post.GetPosts)
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
