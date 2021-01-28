package cmd

import "time"

type Banner struct {
	BusinessId int       `json:"business_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ID         int       `json:"id"`
	Order      int       `json:"order"`
	PicUrl     string    `json:"pic_url"`
	Status     int       `json:"status"`
	StatusStr  string    `json:"status_str"`
	Title      string    `json:"title"`
	Type       string    `json:"type"`
	UserId     int       `json:"user_id"`
}
