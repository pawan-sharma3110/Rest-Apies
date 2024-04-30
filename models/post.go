package models

import "time"

type Post struct {
	ID       int       `json:"id"`
	UserName string    `json:"user_name"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	CreateOn time.Time `json:"created_at"`
	UpdateOn time.Time `json:"updated_at"`
}
