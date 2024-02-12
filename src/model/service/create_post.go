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
	logger.Info("Init createPost model", zap.String("journey", "createPost"))

	postDomainRepository, err := pd.postRepository.CreatePost(postDomain)

	if err != nil {
		return nil, err
	}
	return postDomainRepository, nil
}
