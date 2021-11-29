package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"gravity/gui"
)

var threeBodies = []gui.Body{
	{X: 0, Y: 0, M: 75e8, R: 50, VX: 0, VY: -0.00001},
	{X: 5000, Y: 0, M: 5e8, R: 50, VX: -0.1, VY: 0.5},
	{X: -5000, Y: 0, M: 15e8, R: 50, VX: 0.1, VY: 0.4},
	// {X: 7500, Y:0, M:1e3, R:10, VX:-0.1, VY:0.1},
}

// callback will keep the viewfinder focused on the first body
func callback(ui *gui.UI) {
	bodies := ui.Field.Stats()
	if len(bodies) > 0 {
		central := bodies[0].Position
		v := pixel.V(-float64(central.X), -float64(central.Y))
		ui.Field.ViewFinder.Offset = v
	}
}

func main() {
	ui := gui.NewUI(1024, 1024)
	ui.Load(threeBodies)
	ui.Callback = callback
	ui.Field.ViewFinder.SetScale(30)
	ui.Field.ShowTrails = true
	ui.SetTime(512)
	pixelgl.Run(ui.RunGUI)
}
