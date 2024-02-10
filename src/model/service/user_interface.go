package service

import (
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	// UpdateUser(string) *rest_err.RestErr
	// FindUser(string) (*UserDomain, *rest_err.RestErr)
	// DeleteUser(string) *rest_err.RestErr
}
