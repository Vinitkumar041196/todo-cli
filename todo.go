package main

import (
	"time"
)

type Todo struct {
	Completed bool      `json:"completed"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TodoData struct {
	Todos []Todo
}

func (t *TodoData) add(title string) {
	t.Todos = append(t.Todos, Todo{Title: title, CreatedAt: time.Now()})
}
