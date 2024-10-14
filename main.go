package main

import "log"

func main() {
	data := new(TodoData)
	storage := NewStorage("todo.json")

	err := storage.LoadData(data)
	if err != nil {
		log.Fatal(err)
	}

	err = Execute(data)
	if err != nil {
		log.Fatal(err)
	}

	err = storage.SaveData(data)
	if err != nil {
		log.Fatal(err)
	}
}
