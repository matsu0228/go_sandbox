package main

import (
	"github.com/rivo/tview"
)

func main() {
	// box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	// table := tview.NewTable().SetSelectable(true, false).Select(0, 0).SetFixed(1, 1)
	// table = table.InsertColumn(1)
	// table = table.SetCell(1, 0, tview.NewTableCell("image.ID").
	// 	SetTextColor(tcell.ColorLightYellow).
	// 	SetMaxWidth(1).
	// 	SetExpansion(1))
	//

	sheets := tview.NewList().ShowSecondaryText(false)
	sheets.SetBorder(true).SetTitle("Sheets")
	tables := tview.NewTable().SetBorders(true)
	tables.SetBorder(true).SetTitle("Contents")
	for i, v := range []string{"one", "two", "three"} {
		tables.SetCellSimple(i, 0, v)
	}

	flex := tview.NewFlex().
		AddItem(sheets, 0, 1, true).
		AddItem(tables, 0, 5, false)

	app := tview.NewApplication()
	app = app.SetRoot(flex, true)
	// app = app.SetFocus(table)
	if err := app.Run(); err != nil {
		panic(err)
	}

}
