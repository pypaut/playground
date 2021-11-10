package main

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func table() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	table1 := widgets.NewTable()
	table1.Rows = [][]string{
		[]string{"TOY MACHINES", "ELVES", "SLEIGHS"},
		[]string{"你好吗", "Go-lang is so cool", "Im working on Ruby"},
	}
	table1.TextStyle = ui.NewStyle(ui.ColorWhite)
	tableHeight := 2.5 * float32(len(table1.Rows))
	table1.SetRect(0, 0, 120, int(tableHeight))

	// TEST
	table1.Rows = append(table1.Rows, []string{"HEYBOYZ", "HEYBOYZ", "HEYBOYZ"})
	tableHeight = 2.5 * float32(len(table1.Rows))
	table1.SetRect(0, 0, 120, int(tableHeight))
	// TEST : END

	ui.Render(table1)

	tickerCount := 1
	ticker := time.NewTicker(time.Second / 2).C
	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-ticker:
			table1.Rows[tickerCount][1] = "REPLACED"  // Update a row
			table1.Rows = append(table1.Rows, []string{"HEYBOYZ", "HEYBOYZ", "HEYBOYZ"})  // Add a row
			tableHeight = 2.5 * float32(len(table1.Rows))
			table1.SetRect(0, 0, 120, int(tableHeight))  // Update table rect
			ui.Render(table1)
			tickerCount++
		}
	}
}

func main() {
	table()
	// example()
}
