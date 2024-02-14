package service

import (
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"go.uber.org/zap"
)

func (pd *postDomainService) UpdatePost(postId string, postDomain model.PostDomainInterface) *rest_err.RestErr {
	logger.Info("Init UpdatePost model.",
		zap.String("journey", "UpdatePost"))

	err := pd.postRepository.UpdatePost(postId, postDomain)

	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "UpdatePost"))

		return err
	}

	logger.Info(
		"UpdatePost service executed successfully",
		zap.String("postId", postId),
		zap.String("journey", "UpdatePost"))

	return nil
}
