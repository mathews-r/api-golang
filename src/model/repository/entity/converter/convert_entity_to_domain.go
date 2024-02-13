package converter

import (
	"github.com/mathews-r/golang/src/model"
	"github.com/mathews-r/golang/src/model/repository/entity"
)

func ConvertEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetId(entity.ID.Hex())

	return domain
}

func ConvertEntityToDomainPost(entity entity.PostEntity) model.PostDomainInterface {
	domain := model.NewPostDomain(
		entity.Title,
		entity.Content,
		entity.UserId,
	)

	domain.SetId(entity.ID.Hex())

	return domain
}
