package domain

import "time"

type Todo struct {
	ID        int        `json:"id"`
	Task      string     `json:"task" validate:"required"`
	Done      bool       `json:"done"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
