package dto

import "time"

type NewSocialMediaRequest struct {
	Name           string `json:"name" valid:"required~Name can't be empty" example:"Monday Weeekly Official"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~Social media url can't be empty" example:"https://www.instagram.com/_weeekly/"`
}

type UpdateSocialMediaRequest struct {
	Name           string `json:"name" valid:"required~Name can't be empty" example:"Weeekly Monday Official"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~Social media url can't be empty" example:"https://www.instagram.com/_weeekly/"`
}

type NewSocialMediaResponse struct {
	Id             int       `json:"id" example:"1"`
	Name           string    `json:"name" example:"Monday Weeekly Official"`
	SocialMediaUrl string    `json:"social_media_url" example:"https://www.instagram.com/_weeekly/"`
	UserId         int       `json:"user_id" example:"1"`
	CreatedAt      time.Time `json:"created_at" example:"2023-10-09T05:14:35.19324086+07:00"`
}

type SocialMediaUpdateResponse struct {
	Id             int       `json:"id" example:"1"`
	Name           string    `json:"name" example:"Monday Weeekly Official"`
	SocialMediaUrl string    `json:"social_media_url" example:"https://www.instagram.com/_weeekly/"`
	UserId         int       `json:"user_id" example:"1"`
	UpdatedAt      time.Time `json:"updated_at" example:"2023-10-09T05:14:35.19324086+07:00"`
}

type SocialMediaUser struct {
	Id              int    `json:"id" example:"1"`
	Username        string `json:"username" example:"monday"`
	ProfileImageUrl string `json:"profile_image_url" example:"https://www.pinterest.com/pin/807973989398829161/"`
}

type GetSocialMedia struct {
	Id             int             `json:"id" example:"1"`
	Name           string          `json:"name" example:"Monday Weeekly Official"`
	SocialMediaUrl string          `json:"social_media_url" example:"https://www.instagram.com/_weeekly/"`
	UserId         int             `json:"user_id" example:"1"`
	CreatedAt      time.Time       `json:"created_at" example:"2023-10-09T05:14:35.19324086+07:00"`
	UpdatedAt      time.Time       `json:"updated_at" example:"2023-10-09T05:14:35.19324086+07:00"`
	User           SocialMediaUser `json:"user"`
}

type GetSocialMediaResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

type GetSocialMediaHttpResponse struct {
	StatusCode  int              `json:"status_code" example:"200"`
	Message     string           `json:"message" example:"social media successfully fetched"`
	SocialMedia []*GetSocialMedia `json:"social_media"`
}
