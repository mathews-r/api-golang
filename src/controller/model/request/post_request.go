package request

type PostRequest struct {
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	UserId    int    `json:"userid"`
	Published string `json:"published"`
	Updated   string `json:"updated"`
}
