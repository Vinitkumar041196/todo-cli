package main

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
)

func FormatToTable(data [][]interface{}) {
	w := table.NewWriter()
	w.AppendHeader(data[0])
	for _, row := range data[1:] {
		w.AppendRow(row)
	}
	w.SetStyle(table.StyleColoredBlackOnGreenWhite)
	str := w.Render()
	fmt.Println(str)
}
