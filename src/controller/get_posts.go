package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/view"
	"go.uber.org/zap"
)

func (pc *postControllerInterface) GetPosts(c *gin.Context) {
	logger.Info("Init GetPost controller",
		zap.String("journey", "GetPost"),
	)

	posts, err := pc.service.GetPosts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (pc *postControllerInterface) GetPostById(c *gin.Context) {
	postId := c.Param("postId")
	post, err := pc.service.GetPostById(postId)

	if err != nil {
		logger.Info("Error trying to call GetPost service",
			zap.String("journey", "GetPost"),
		)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	logger.Info(
		"CreatePost controller executed successfully",
		zap.String("postId", postId),
		zap.String("journey", "GetPost"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponsePost(post))
}
