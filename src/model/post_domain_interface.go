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
	GetUserId() int
	SetId(string)
}

func NewPostDomain(
	title, content string, userId int,
) PostDomainInterface {
	return &postDomain{
		Title:     title,
		Content:   content,
		UserId:    userId,
		Published: time.Now(),
		Updated:   time.Now(),
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
