package response

import "time"

type PostResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserId    int       `json:"userid"`
	Published time.Time `json:"published"`
	Updated   time.Time `json:"updated"`
}
