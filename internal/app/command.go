package app

import (
	"flag"
	"fmt"
	"os"
)

func parseSubFlags() (int, string, bool, error) {
	var index int
	var title string
	var autoPrint bool
	subcommands := flag.NewFlagSet("", flag.ExitOnError)
	subcommands.IntVar(&index, "index", -1, "Index of todo")
	subcommands.IntVar(&index, "i", -1, "Index of todo")
	subcommands.StringVar(&title, "title", "", "New title for todo as a string")
	subcommands.StringVar(&title, "t", "", "New title for todo as a string")
	subcommands.BoolVar(&autoPrint, "o", false, "Automatically print latest todo table")
	err := subcommands.Parse(os.Args[2:])
	return index, title, autoPrint, err
}

func (a *App) Execute() error {
	index, title, autoPrint, err := parseSubFlags()
	if err != nil {
		return err
	}

	a.setAutoPrint(autoPrint)

	switch os.Args[1] {
	case "add":
		return a.add(title)
	case "edit":
		return a.edit(index-1, title)
	case "delete":
		return a.delete(index - 1)
	case "toggle":
		return a.toggleCompletion(index - 1)
	case "list":
		return a.list()
	default:
		return fmt.Errorf("invalid command %v", os.Args[1])
	}
}
