package main

import (
	"flag"
	"fmt"
	"os"
)

func parseSubFlags() (int, string, error) {
	var index int
	var title string
	subcommands := flag.NewFlagSet("", flag.ExitOnError)
	subcommands.IntVar(&index, "index", -1, "Index of todo")
	subcommands.StringVar(&title, "title", "", "New title for todo as a string")
	err := subcommands.Parse(os.Args[2:])
	return index, title, err
}

func Execute(data *TodoData) error {
	switch os.Args[1] {
	case "add":
		_, title, err := parseSubFlags()
		if err != nil {
			return err
		}
		return data.add(title)
	case "edit":
		index, title, err := parseSubFlags()
		if err != nil {
			return err
		}
		return data.edit(index-1, title)
	case "delete":
		index, _, err := parseSubFlags()
		if err != nil {
			return err
		}
		return data.delete(index - 1)
	case "toggle":
		index, _, err := parseSubFlags()
		if err != nil {
			return err
		}
		return data.toggleCompletion(index - 1)
	case "list":
		return data.list()
	default:
		return fmt.Errorf("invalid command %v", os.Args[1])
	}
}
