package request

import "time"

type PostRequest struct {
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	UserId    int       `json:"userid" binding:"required"`
	Published time.Time `json:"published"`
	Updated   time.Time `json:"updated"`
}
