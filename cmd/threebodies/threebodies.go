package main

import (
	"github.com/clambin/gravity/gui"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk/vect"
	"golang.org/x/image/colornames"
)

const position = 1500
const offset = 125

var threeBodies = []gui.Body{
	{Position: vect.Vect{X: 0, Y: 0}, Mass: 3e7, Radius: 150, Color: colornames.Yellow},
	{Position: vect.Vect{X: -position, Y: position}, Mass: 2.0e6, Radius: 100, Velocity: vect.Vect{X: 0, Y: offset}, Color: colornames.Red},
	{Position: vect.Vect{X: position, Y: -position}, Mass: 2.4e6, Radius: 105, Velocity: vect.Vect{X: 0, Y: -offset}, Color: colornames.Green},
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
