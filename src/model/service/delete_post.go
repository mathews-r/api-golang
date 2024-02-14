package service

import (
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"go.uber.org/zap"
)

func (pd *postDomainService) DeletePost(postId string) *rest_err.RestErr {
	logger.Info("Init DeletePost model.",
		zap.String("journey", "DeletePost"))

	err := pd.postRepository.DeletePost(postId)

	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "DeletePost"))

		return err
	}

	logger.Info(
		"DeletePost service executed successfully",
		zap.String("postId", postId),
		zap.String("journey", "DeletePost"))

	return nil
}
