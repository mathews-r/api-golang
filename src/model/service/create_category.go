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
	logger.Info("Init CreateCategory model", zap.String("journey", "CreateCategory"))

	categoryDomainRepository, err := cd.categoryRepository.CreateCategory(categoryDomain)

	if err != nil {
		return nil, err
	}
	return categoryDomainRepository, nil
}
