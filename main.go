package main

import "log"

func main() {
	data := new(TodoData)
	err := Execute(data)
	if err != nil {
		log.Fatal(err)
	}
}
