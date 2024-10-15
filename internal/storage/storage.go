package storage

import "github.com/Vinitkumar041196/todo-cli/internal/types"

type Storage interface {
	LoadData() ([]types.Todo, error)
	SaveData(data []types.Todo) error
}
