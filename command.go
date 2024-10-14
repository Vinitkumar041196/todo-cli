package main

import (
	"flag"
	"fmt"
	"os"
)

func parseSubFlags() (int, string) {
	var index int
	var title string
	subcommands := flag.NewFlagSet("modify", flag.ExitOnError)
	subcommands.IntVar(&index, "index", -1, "Index of todo")
	subcommands.StringVar(&title, "title", "", "New title for todo as a string")
	subcommands.Parse(os.Args[2:])
	return index, title
}

func Execute(data *TodoData) error {
	switch os.Args[1] {
	case "add":
		_, title := parseSubFlags()
		data.add(title)
		return nil
	case "edit":
		index, title := parseSubFlags()
		return data.edit(index-1, title)
	case "delete":
		index, _ := parseSubFlags()
		return data.delete(index - 1)
	case "toggle":
		index, _ := parseSubFlags()
		return data.toggleCompletion(index - 1)
	case "list":
		return data.list()
	default:
		return fmt.Errorf("invalid command %v", os.Args[1])
	}
}
