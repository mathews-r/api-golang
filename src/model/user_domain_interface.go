package model

import "github.com/mathews-r/golang/src/configs/rest_err"

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int
	EncryptPassword()
	SetId(string)
	GetId() string
	GenerateToken() (string, *rest_err.RestErr)
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

func LoginDomain(
	email, password string,
) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
	}
}

func UpdateUserDomain(
	name string,
	age int,
) UserDomainInterface {
	return &userDomain{
		Name: name,
		Age:  age,
	}
}
