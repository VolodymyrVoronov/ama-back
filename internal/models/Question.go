package models

import "time"

type Question struct {
	ID          int       `json:"id"`
	AuthorEmail string    `json:"author_email"`
	Question    string    `json:"question"`
	Answer      string    `json:"answer"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
