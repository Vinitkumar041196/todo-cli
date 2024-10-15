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

func (t *TodoData) validateIndex(index int) error {
	if index < 0 || index >= len(t.Todos) {
		return fmt.Errorf("invalid index")
	}
	return nil
}

func validateTitle(title string) error {
	if title == "" {
		return fmt.Errorf("invalid title")
	}
	return nil
}
func (t *TodoData) add(title string) error {
	err := validateTitle(title)
	if err != nil {
		return err
	}

	t.Todos = append(t.Todos, Todo{Title: title, CreatedAt: time.Now()})
	return nil
}

func (t *TodoData) edit(index int, title string) error {
	err := t.validateIndex(index)
	if err != nil {
		return err
	}

	err = validateTitle(title)
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

func (t *TodoData) list() error {
	data := [][]interface{}{
		{"#", "Status", "Title", "Created At", "Updated At"},
	}

	completedCount := 0
	for i, todo := range t.Todos {

		completed := "❌"
		if todo.Completed {
			completed = "✅"
			completedCount++
		}

		updatedAt := "-"
		if !todo.UpdatedAt.IsZero() {
			updatedAt = todo.UpdatedAt.Format("2006-01-02 03:04:05PM")
		}

		data = append(data, []interface{}{fmt.Sprint(i + 1), completed, todo.Title, todo.CreatedAt.Format("2006-01-02 03:04:05PM"), updatedAt})
	}

	data = append(data, []interface{}{"", "", "", "Tasks Completed", fmt.Sprintf("%d/%d", completedCount, len(t.Todos))})

	outStr := formatToTable(data)

	fmt.Println("\n\n" + outStr + "\n\n")

	return nil
}
