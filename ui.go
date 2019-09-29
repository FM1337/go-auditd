package main

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func initUI() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
}

func renderList() *widgets.List {
	l := widgets.NewList()
	l.Title = "AuditD Logs"
	l.Rows = testPrint()
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = true
	x2, _ := ui.TerminalDimensions()
	l.SetRect(0, 0, x2, 20)

	ui.Render(l)

	return l
}

func renderTable() *widgets.Table {
	countTypes()
	table := widgets.NewTable()
	table.Title = "AuditD Stats"
	table.Rows = append(table.Rows, []string{
		"Type",
		"Count",
	})
	for key, value := range types {
		tmpSlice := []string{}
		tmpSlice = append(tmpSlice, key)
		tmpSlice = append(tmpSlice, fmt.Sprintf("%d", value))
		table.Rows = append(table.Rows, tmpSlice)
	}
	x2, y2 := ui.TerminalDimensions()
	table.SetRect(0, 20, x2, y2)
	return table
}

func renderElements() {
	l := renderList()
	table := renderTable()

	ui.Render(table, l)

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			ui.Close()
			return
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "<C-d>":
			l.ScrollHalfPageDown()
		case "<C-u>":
			l.ScrollHalfPageUp()
		case "<C-f>":
			l.ScrollPageDown()
		case "<C-b>":
			l.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				l.ScrollTop()
			}
		case "<Home>":
			l.ScrollTop()
		case "G", "<End>":
			l.ScrollBottom()
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		ui.Render(l)
	}
}
