package converter

import (
	"github.com/mathews-r/golang/src/model"
	"github.com/mathews-r/golang/src/model/repository/entity"
)

func ConvertDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Name:     domain.GetName(),
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Age:      domain.GetAge(),
	}
}

func ConvertDomainToEntityPost(domain model.PostDomainInterface) *entity.PostEntity {
	return &entity.PostEntity{
		Title:     domain.GetTitle(),
		Content:   domain.GetContent(),
		UserId:    domain.GetUserId(),
		Published: domain.GetPublished(),
		Updated:   domain.GetUpdated(),
	}
}
