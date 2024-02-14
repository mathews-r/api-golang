package request

type CategoryRequest struct {
	ID       string `json:"id"`
	Category string `json:"category" binding:"required"`
}
