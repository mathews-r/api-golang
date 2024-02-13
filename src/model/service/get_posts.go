package service

import (
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"go.uber.org/zap"
)

func (pd *postDomainService) GetPosts() ([]model.PostDomainInterface, *rest_err.RestErr) {

	logger.Info("Init GetPosts model", zap.String("journey", "GetPosts"))

	postDomainRepository, err := pd.postRepository.GetPosts()

	if err != nil {
		return nil, err
	}
	return postDomainRepository, nil
}
