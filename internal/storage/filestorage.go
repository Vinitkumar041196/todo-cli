package storage

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Vinitkumar041196/todo-cli/internal/types"
)

type FileStorage struct {
	fileName string
}

func NewStorage(fileName string) Storage {
	f, err := os.Open(fileName)
	f.Close()

	if err != nil {
		f, _ := os.Create(fileName)
		f.Close()
	}

	return &FileStorage{fileName: fileName}
}

func (s *FileStorage) LoadData() ([]types.Todo, error) {
	byt, err := os.ReadFile(s.fileName)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	data := []types.Todo{}
	if len(byt) > 0 {
		err = json.Unmarshal(byt, &data)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return data, nil
}

func (s *FileStorage) SaveData(data []types.Todo) error {
	byt, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		log.Println(err)
		return err
	}

	err = os.WriteFile(s.fileName, byt, 0644)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
