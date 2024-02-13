package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (pc *postControllerInterface) GetPosts(c *gin.Context) {
	posts, err := pc.service.GetPosts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, posts)
}
