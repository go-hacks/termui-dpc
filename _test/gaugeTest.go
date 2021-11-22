package main

import (
	"log"
	"time"
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	colors := make([]ui.Color, 6)
	colors[0] = ui.ColorGreen
	colors[1] = ui.ColorBlue
	colors[2] = ui.ColorMagenta
	colors[3] = ui.ColorCyan
	colors[4] = ui.ColorWhite
	colors[5] = ui.ColorRed

	g := widgets.NewGauge()
	g.Title = "Gauge"
	g.SetRect(0, 12, 50, 15)
	g.BarColor = ui.ColorBlue
	g.BorderStyle.Fg = ui.ColorWhite
	g.TitleStyle.Fg = ui.ColorCyan
	g.LabelStyle.Fg = colors[0]

	tickerCount := 1
	ui.Render(g)
	tickerCount++
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(50 * time.Millisecond).C
	colorCnt := 0
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
				case "q", "<C-c>":
					return
			}
		case <-ticker:
			ui.Render(g)
			g.Percent = tickerCount % 101
			if tickerCount % 101 == 0 {
				colorCnt++
				if colorCnt == 6 {
					colorCnt = 0
				}
				g.LabelStyle.Fg = colors[colorCnt]
			}
			tickerCount++
		}
	}

	return
}
