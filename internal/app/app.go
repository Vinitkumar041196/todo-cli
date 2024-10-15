package app

import (
	"fmt"
	"time"

	"github.com/Vinitkumar041196/todo-cli/internal/storage"
	"github.com/Vinitkumar041196/todo-cli/internal/types"
)

type App struct {
	todos     []types.Todo
	autoPrint bool
	store     storage.Storage
}

func NewApp(store storage.Storage) *App {
	return &App{todos: make([]types.Todo, 0), store: store}
}

func (a *App) LoadData() error {
	var err error
	a.todos, err = a.store.LoadData()
	return err
}

func (a *App) SaveData() error {
	return a.store.SaveData(a.todos)
}

func (a *App) validateIndex(index int) error {
	if index < 0 || index >= len(a.todos) {
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

func (a *App) setAutoPrint(val bool) {
	a.autoPrint = val
}

func (a *App) add(title string) error {
	err := validateTitle(title)
	if err != nil {
		return err
	}

	a.todos = append(a.todos, types.Todo{Title: title, CreatedAt: time.Now()})

	if a.autoPrint {
		return a.list()
	}

	return nil
}

func (a *App) edit(index int, title string) error {
	err := a.validateIndex(index)
	if err != nil {
		return err
	}

	err = validateTitle(title)
	if err != nil {
		return err
	}

	todo := a.todos[index]
	todo.Title = title
	todo.UpdatedAt = time.Now()
	a.todos[index] = todo

	if a.autoPrint {
		return a.list()
	}

	return nil
}

func (a *App) toggleCompletion(index int) error {
	err := a.validateIndex(index)
	if err != nil {
		return err
	}

	todo := a.todos[index]
	todo.Completed = !todo.Completed
	todo.UpdatedAt = time.Now()
	a.todos[index] = todo

	if a.autoPrint {
		return a.list()
	}

	return nil
}

func (a *App) delete(index int) error {
	err := a.validateIndex(index)
	if err != nil {
		return err
	}

	a.todos = append(a.todos[:index], a.todos[index+1:]...)

	if a.autoPrint {
		return a.list()
	}

	return nil
}

func (a *App) list() error {
	data := [][]interface{}{
		{"#", "Status", "Title", "Created At", "Updated At"},
	}

	completedCount := 0
	for i, todo := range a.todos {

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

	data = append(data, []interface{}{"", "", "", "Tasks Completed", fmt.Sprintf("%d/%d", completedCount, len(a.todos))})

	outStr := formatToTable(data)

	fmt.Println("\n\n" + outStr + "\n\n")

	return nil
}
