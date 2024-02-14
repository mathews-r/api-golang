package service

import (
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"go.uber.org/zap"
)

func (cd *categoryDomainService) CreateCategory(
	categoryDomain model.CategoryDomainInterface,
) (model.CategoryDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateCategory model.",
		zap.String("journey", "CreateCategory"))

	categoryDomainRepository, err := cd.categoryRepository.CreateCategory(categoryDomain)

	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "CreateCategory"))
		return nil, err
	}

	logger.Info(
		"CreateCategory service executed successfully",
		zap.String("categoryId", categoryDomainRepository.GetId()),
		zap.String("journey", "CreateCategory"))

	return categoryDomainRepository, nil
}
