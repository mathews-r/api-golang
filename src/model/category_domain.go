package model

type categoryDomain struct {
	ID       string
	Category string
}

func (cd *categoryDomain) GetId() string {
	return cd.ID
}

func (cd *categoryDomain) GetCategory() string {
	return cd.Category
}

func (cd *categoryDomain) SetId(id string) {
	cd.ID = id
}
