package main

func main() {
	data := new(TodoData)
	data.add("Task1")
	data.add("Task2")
	data.list()
	data.edit(1, "Task 2 updated")
	data.toggleCompletion(0)
	data.list()
	data.delete(1)
	data.add("Task3")
	data.list()
}
