package main

import (
	"log"

	"github.com/Vinitkumar041196/todo-cli/internal/app"
	"github.com/Vinitkumar041196/todo-cli/internal/storage"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate)

	storage := storage.NewStorage("todo.json")
	app := app.NewApp(storage)

	err := app.LoadData()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Execute()
	if err != nil {
		log.Fatal(err)
	}

	err = app.SaveData()
	if err != nil {
		log.Fatal(err)
	}
}
