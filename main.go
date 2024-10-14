package main

import "log"

func main() {
	data := new(TodoData)
	storage := NewStorage("todo.json")

	err := Execute(data)
	if err != nil {
		log.Fatal(err)
	}
}
