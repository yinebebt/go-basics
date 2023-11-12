package main

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func main() {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "Name", "Score", "Added")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	// for _, widget := range getWidgets() {
	// 	tbl.AddRow(widget.ID, widget.Name, widget.Cost, widget.Added)
	// }

	tbl.Print()
}
