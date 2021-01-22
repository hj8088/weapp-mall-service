package cmd

import "time"

type Banner struct {
	BusinessId int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ID         int
	Order      int
	PicUrl     string
	Status     int
	StatusStr  string
	Title      string
	Type       string
	UserId     int
}
