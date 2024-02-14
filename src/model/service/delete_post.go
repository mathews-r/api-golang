package service

import (
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"go.uber.org/zap"
)

func (pd *postDomainService) DeletePost(postId string) *rest_err.RestErr {
	logger.Info("Init DeletePost model", zap.String("journey", "DeletePost"))

	err := pd.postRepository.DeletePost(postId)

	if err != nil {
		return err
	}
	return nil
}
