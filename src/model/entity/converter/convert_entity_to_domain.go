package converter

import (
	"github.com/mathews-r/golang/src/model"
	"github.com/mathews-r/golang/src/model/entity"
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
