package main

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

func formatToTable(data [][]interface{}) string {
	w := table.NewWriter()
	w.AppendHeader(data[0])
	for _, row := range data[1 : len(data)-1] {
		w.AppendRow(row)
	}
	w.AppendFooter(data[len(data)-1])
	w.SetStyle(table.StyleColoredBlackOnGreenWhite)
	return w.Render()
}
