package main

import (
	"github.com/clambin/gravity/gui"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk/vect"
)

var singleSun = []gui.Body{
	{Position: vect.Vect{X: 0, Y: 0}, Mass: 3.33e7, Radius: 50, Velocity: vect.Vect{X: 0, Y: 0}},
	{Position: vect.Vect{X: 500, Y: 0}, Mass: 1, Radius: 10, Velocity: vect.Vect{X: 0, Y: 40}},
	{Position: vect.Vect{X: 1000, Y: 0}, Mass: 1, Radius: 10, Velocity: vect.Vect{X: 0, Y: 40}},
	{Position: vect.Vect{X: 1500, Y: 0}, Mass: 1e2, Radius: 25, Velocity: vect.Vect{X: 0, Y: 40}},
}

func main() {
	ui := gui.NewUI("single sun", 1024, 1024)
	ui.Field.ShowTrails = true
	ui.Load(singleSun)
	ui.Field.ViewFinder.SetScale(5)
	pixelgl.Run(ui.RunGUI)
}
