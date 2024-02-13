package service

import (
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("Init LoginUser model", zap.String("journey", "LoginUser"))

	userDomain.EncryptPassword()
	user, err := ud.findUserByEmailAndPassword(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
