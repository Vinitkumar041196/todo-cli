package main

import (
	"fmt"
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

func (t *TodoData) validateIndex(index int) error {
	if index < 0 || index >= len(t.Todos) {
		return fmt.Errorf("invalid index")
	}
	return nil
}

func (t *TodoData) edit(index int, title string) error {
	err := t.validateIndex(index)
	if err != nil {
		return err
	}

	todo := t.Todos[index]
	todo.Title = title
	todo.UpdatedAt = time.Now()
	t.Todos[index] = todo

	return nil
}

func (t *TodoData) toggleCompletion(index int) error {
	err := t.validateIndex(index)
	if err != nil {
		return err
	}

	todo := t.Todos[index]
	todo.Completed = !todo.Completed
	todo.UpdatedAt = time.Now()
	t.Todos[index] = todo

	return nil
}

func (t *TodoData) delete(index int) error {
	err := t.validateIndex(index)
	if err != nil {
		return err
	}

	t.Todos = append(t.Todos[:index], t.Todos[index+1:]...)
	return nil
}
