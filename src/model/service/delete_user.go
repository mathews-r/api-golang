package service

import (
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("Init DeleteUser model", zap.String("journey", "DeleteUser"))

	err := ud.userRepository.DeleteUser(userId)

	if err != nil {
		return err
	}
	return nil
}
