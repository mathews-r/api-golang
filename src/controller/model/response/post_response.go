package response

type PostResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Category  string `json:"category"`
	UserId    string `json:"userid"`
	Published string `json:"published"`
	Updated   string `json:"updated"`
}
