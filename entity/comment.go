package entity

import "time"

type Comment struct {
	Id        int
	UserId    int
	PhotoId   int
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
