package model

import (
	"encoding/json"
	"fmt"
)

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

func (ud *userDomain) GetJSONValue() (string, error) {
	jsonValues, err := json.Marshal(ud)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(jsonValues), nil
}
