package model

type userDomain struct {
	ID       string
	Email    string
	Password string
	Name     string
	Age      int
}

func (ud *userDomain) GetId() string {
	return ud.ID
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
