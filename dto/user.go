package dto

import "time"

type NewUserRequest struct {
	Username string `json:"username" valid:"required~Username can't be empty" example:"monday"`
	Email    string `json:"email" valid:"required~Email can't be empty, email~Email has to be a valid email format" example:"monday.day@email.com"`
	Age      int    `json:"age" valid:"required~Age can't be empty, range(9|150)~Minimum age requirement is 9 years old" example:"21"`
	Password string `json:"password" valid:"required~Password can't be empty, length(6|255)~Minimum password characters are 6 characters" example:"secret"`
}

type UserLoginRequest struct {
	Email    string `json:"email" valid:"required~Email can't be empty, email~Email has to be a valid email format" example:"monday.day@email.com"`
	Password string `json:"password" valid:"required~Password can't be empty" example:"secret"`
}

type UserUpdateRequest struct {
	Username string `json:"username" valid:"required~Username can't be empty" example:"monday"`
	Email    string `json:"email" valid:"required~Email can't be empty, email~Email has to be a valid email format" example:"monday.day@weeekly.com"`
}

type UserUpdateResponse struct {
	Id        int       `json:"id" example:"1"`
	Username  string    `json:"username" example:"monday"`
	Email     string    `json:"email" example:"monday.day@weeekly.com"`
	Age       uint      `json:"age" example:"21"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-10-09T05:14:35.19324086+07:00"`
}

type UserResponse struct {
	Id       int    `json:"id" example:"1"`
	Username string `json:"username" example:"monday"`
	Email    string `json:"email" example:"monday.day@email.com"`
	Age      uint   `json:"age" example:"21"`
}

type TokenResponse struct {
	Token string `json:"token" example:"random string"`
}

type GetUserResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
