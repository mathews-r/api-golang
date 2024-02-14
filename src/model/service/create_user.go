package service

import (
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser model.",
		zap.String("journey", "createUser"))

	user, _ := ud.FindUserByEmail(userDomain.GetEmail())
	if user != nil {
		return nil, rest_err.NewBadRequestErr("User already exists")
	}

	userDomain.EncryptPassword()
	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)

	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info(
		"CreateUser service executed successfully",
		zap.String("userId", userDomainRepository.GetId()),
		zap.String("journey", "createUser"))
	return userDomainRepository, nil
}
