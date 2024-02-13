package model

type postDomain struct {
	ID        string
	Title     string
	Content   string
	UserId    int
	Published string
	Updated   string
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

func (pd *postDomain) GetPublished() string {
	return pd.Published
}

func (pd *postDomain) GetUpdated() string {
	return pd.Updated
}
