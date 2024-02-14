package service

import (
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"go.uber.org/zap"
)

func (pd *postDomainService) CreatePost(
	postDomain model.PostDomainInterface,
) (model.PostDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreatePost model.",
		zap.String("journey", "CreatePost"))

	postDomainRepository, err := pd.postRepository.CreatePost(postDomain)

	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "CreatePost"))

		return nil, err
	}

	logger.Info(
		"CreatePost service executed successfully",
		zap.String("categoryId", postDomainRepository.GetCategory()),
		zap.String("journey", "CreatePost"))

	return postDomainRepository, nil
}
