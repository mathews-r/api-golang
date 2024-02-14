package model

type CategoryDomainInterface interface {
	GetId() string
	GetCategory() string
	SetId(string)
}

func CategoryDomain(
	category string,
) CategoryDomainInterface {
	return &categoryDomain{
		Category: category,
	}
}
