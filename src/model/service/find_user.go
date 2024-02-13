package service

import (
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail", zap.String("journey", "findUserByEmail"))
	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) findUserByEmailAndPassword(email string, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmailAndPassword", zap.String("journey", "findUserByEmailAndPassword"))
	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
