package entity

import "time"

type SocialMedia struct {
	Id             int
	Name           string
	SocialMediaUrl string
	UserId         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
