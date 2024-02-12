package converter

import (
	"github.com/mathews-r/golang/src/model"
	"github.com/mathews-r/golang/src/model/entity"
)

func ConvertDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Name:     domain.GetName(),
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
	}
}
