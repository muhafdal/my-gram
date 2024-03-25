package entity

import "time"

type Photo struct {
	Id        int
	Title     string
	Caption   string
	PhotoUrl  string
	UserId    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
