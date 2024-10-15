package main

import (
	"encoding/json"
	"log"
	"os"
)

type Storage struct {
	fileName string
}

func NewStorage(fileName string) *Storage {
	f, err := os.Open(fileName)
	f.Close()

	if err != nil {
		f, _ := os.Create(fileName)
		f.Close()
	}

	return &Storage{fileName: fileName}
}

func (s *Storage) LoadData(data *TodoData) error {
	byt, err := os.ReadFile(s.fileName)
	if err != nil {
		log.Println(err)
		return err
	}

	if len(byt) > 0 {
		err = json.Unmarshal(byt, &data.Todos)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (s *Storage) SaveData(data *TodoData) error {
	byt, err := json.MarshalIndent(data.Todos, "", "	")
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
