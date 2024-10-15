package types

import (
	"time"
)

type Todo struct {
	Completed bool      `json:"completed"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
