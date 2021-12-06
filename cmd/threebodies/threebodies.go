package main

import (
	"github.com/clambin/gravity/gui"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk/vect"
)

var threeBodies = []gui.Body{
	{Position: vect.Vect{X: 0, Y: 0}, Mass: 20e6, Radius: 100},
	{Position: vect.Vect{X: -2500, Y: 2500}, Mass: 2.0e6, Radius: 50, Velocity: vect.Vect{Y: -25}},
	{Position: vect.Vect{X: 2500, Y: -2500}, Mass: 2.4e6, Radius: 55, Velocity: vect.Vect{Y: 25}},
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
	ui := gui.NewUI("three bodies", 1024, 1024)
	ui.Load(threeBodies)
	ui.Callback = callback
	ui.Field.ViewFinder.SetScale(10)
	ui.Field.ShowTrails = true
	ui.SetTime(64)
	pixelgl.Run(ui.RunGUI)
}
