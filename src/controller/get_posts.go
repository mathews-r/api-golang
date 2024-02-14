package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/view"
)

func (pc *postControllerInterface) GetPosts(c *gin.Context) {
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
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, view.ConvertDomainToResponsePost(post))
}
