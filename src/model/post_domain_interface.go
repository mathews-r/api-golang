package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type PostDomainInterface interface {
	GetPostId() string
	GetTitle() string
	GetContent() string
	GetUserId() string
	SetId(string)
	GetPublished() string
	GetUpdated() string
	GetCategory() string
}

func (pd *postDomain) GetJSONValue() (string, error) {

	jsonValues, err := json.Marshal(pd)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(jsonValues), nil
}

func NewPostDomain(
	title, content, category, userId, published, updated string,
) PostDomainInterface {
	return &postDomain{
		Title:     title,
		Content:   content,
		Category:  category,
		UserId:    userId,
		Published: time.DateTime,
		Updated:   time.DateTime,
	}
}

func UpdatePostDomain(
	title, category, content string,
) PostDomainInterface {
	return &postDomain{
		Title:    title,
		Content:  content,
		Category: category,
	}
}
