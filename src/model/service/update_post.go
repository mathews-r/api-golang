package service

import (
	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"github.com/mathews-r/golang/src/model"
	"go.uber.org/zap"
)

func (pd *postDomainService) UpdatePost(postId string, postDomain model.PostDomainInterface) *rest_err.RestErr {
	logger.Info("Init updatePost model", zap.String("journey", "updatePost"))

	err := pd.postRepository.UpdatePost(postId, postDomain)

	if err != nil {
		return err
	}
	return nil
}
