package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int
	EncryptPassword()
	SetId(string)
	GetId() string
}

func NewUserDomain(
	email, password, name string,
	age int,
) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}
