package service

import (
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"go.uber.org/zap"
)

func (pd *postDomainService) GetPosts() ([]model.PostDomainInterface, *rest_err.RestErr) {

	logger.Info("Init GetPosts model.",
		zap.String("journey", "GetPosts"))

	postDomainRepository, err := pd.postRepository.GetPosts()

	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "GetPosts"))

		return nil, err
	}

	logger.Info(
		"GetPosts service executed successfully",
		zap.String("journey", "GetPosts"))
	return postDomainRepository, nil
}

func (pd *postDomainService) GetPostById(postId string) (model.PostDomainInterface, *rest_err.RestErr) {

	logger.Info("Init GetPostById model.",
		zap.String("journey", "GetPostById"))

	postDomainRepository, err := pd.postRepository.GetPostById(postId)

	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "GetPostById"))

		return nil, err
	}

	logger.Info(
		"GetPostById service executed successfully",
		zap.String("postId", postDomainRepository.GetPostId()),
		zap.String("journey", "GetPostById"))

	return postDomainRepository, nil
}
