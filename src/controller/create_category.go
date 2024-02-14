package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/validation"
	"github.com/mathews-r/golang/src/controller/model/request"
	"github.com/mathews-r/golang/src/model"
	"github.com/mathews-r/golang/src/view"
	"go.uber.org/zap"
)

var (
	CategoryDomainInterface model.CategoryDomainInterface
)

func (cc *categoryControllerInterface) CreateCategory(c *gin.Context) {
	var categoryRequest request.CategoryRequest

	if err := c.ShouldBindJSON(&categoryRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
	}

	domain := model.CategoryDomain(
		categoryRequest.Category,
	)

	domainResult, err := cc.service.CreateCategory(domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("category created successfully", zap.String("journey", "categoryCreate"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponseCategory(domainResult))
}
