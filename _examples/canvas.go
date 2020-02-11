// +build ignore

package main

import (
	"image"
	"log"

	ui "github.com/tim-koehler/termui/v3"
)

// TODO make example interactive so you can draw lines
func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	c := ui.NewCanvas()
	c.SetRect(0, 0, 50, 50)
	c.SetLine(image.Pt(0, 0), image.Pt(10, 20), ui.ColorWhite)

	ui.Render(c)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
