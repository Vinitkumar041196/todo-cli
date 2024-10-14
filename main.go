package main

import "time"

type TodoData struct {
	Todos []Todo
}

type Todo struct {
	Completed bool
	Title     string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func main() {

}
