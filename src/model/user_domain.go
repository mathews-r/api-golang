package model

import (
	"encoding/json"
	"fmt"
)

type userDomain struct {
	ID       string
	Email    string
	Password string
	Name     string
	Age      int
}

func (ud *userDomain) GetJSONValue() (string, error) {
	jsonValues, err := json.Marshal(ud)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(jsonValues), nil
}

func (ud *userDomain) SetId(id string) {
	ud.ID = id
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}

func (ud *userDomain) GetPassword() string {
	return ud.Password
}

func (ud *userDomain) GetName() string {
	return ud.Name
}

func (ud *userDomain) GetAge() int {
	return ud.Age
}
