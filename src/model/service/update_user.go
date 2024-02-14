package service

import (
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateUser model.",
		zap.String("journey", "updateUser"))

	err := ud.userRepository.UpdateUser(userId, userDomain)

	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "updateUser"))

		return err
	}

	logger.Info(
		"updateUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))
	return nil
}
