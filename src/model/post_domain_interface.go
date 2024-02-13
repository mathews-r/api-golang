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
}

func NewPostDomain(
	title, content, published, updated, userId string,
) PostDomainInterface {
	return &postDomain{
		Title:     title,
		Content:   content,
		Published: time.DateTime,
		Updated:   time.DateTime,
		UserId:    userId,
	}
}

func (pd *postDomain) GetJSONValue() (string, error) {

	jsonValues, err := json.Marshal(pd)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(jsonValues), nil
}
