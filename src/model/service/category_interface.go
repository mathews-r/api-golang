package service

import (
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"github.com/mathews-r/golang/src/model/repository"
)

func NewCategoryDomainService(categoryInterface repository.CategoryRepository) CategoryDomainService {
	return &categoryDomainService{categoryInterface}
}

type categoryDomainService struct {
	categoryRepository repository.CategoryRepository
}

type CategoryDomainService interface {
	CreateCategory(model.CategoryDomainInterface) (model.CategoryDomainInterface, *rest_err.RestErr)
}
