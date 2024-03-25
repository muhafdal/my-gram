package dto

import (
	"time"
)

type NewCommentRequest struct {
	PhotoId int    `json:"photo_id" example:"1"`
	Message string `json:"message" valid:"required~Message can't be empty" example:"so beautiful"`
}

type NewCommentResponse struct {
	Id        int       `json:"id" example:"1"`
	UserId    int       `json:"user_id" example:"1"`
	PhotoId   int       `json:"photo_id" example:"1"`
	Message   string    `json:"message" example:"so beautifull"`
	CreatedAt time.Time `json:"created_at" example:"2023-10-09T05:14:35.19324086+07:00"`
}

type GetCommentResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

type UpdateCommentRequest struct {
	Message string `json:"message" valid:"required~Message can't be empty" example:"omg so beautiful"`
}

type UpdateCommentResponseData struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	PhotoId   int       `json:"photo_id"`
	Message   string    `json:"message"`
	UpdatedAt time.Time `json:"updated_at"`
}
type UpdateCommentResponse struct {
	StatusCode int                       `json:"status_code"`
	Message    string                    `json:"message"`
	Data       UpdateCommentResponseData `json:"data"`
}
