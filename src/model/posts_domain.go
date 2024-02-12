package model

import (
	"time"
)

type postDomain struct {
	ID        string
	Title     string
	Content   string
	UserId    int
	Published time.Time
	Updated   time.Time
}

func (pd *postDomain) SetId(id string) {
	pd.ID = id
}

func (pd *postDomain) GetTitle() string {
	return pd.Title
}

func (pd *postDomain) GetContent() string {
	return pd.Content
}

func (pd *postDomain) GetUserId() int {
	return pd.UserId
}

func (pd *postDomain) GetPostId() string {
	return pd.ID
}

func (pd *postDomain) GetPublised() string {
	return pd.Published.String()
}

func (pd *postDomain) GetUpdate() string {
	return pd.Updated.String()
}
