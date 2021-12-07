package main

import (
	"github.com/clambin/gravity/gui"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk/vect"
	"golang.org/x/image/colornames"
)

var singleSun = []gui.Body{
	{Position: vect.Vect{X: 0, Y: 0}, Mass: 3.33e7, Radius: 50, Velocity: vect.Vect{X: 0, Y: 0}, Color: colornames.Yellow},
	{Position: vect.Vect{X: 200, Y: 0}, Mass: 1e2, Radius: 10, Velocity: vect.Vect{X: 0, Y: 440}, Color: colornames.Grey},
	{Position: vect.Vect{X: 600, Y: 0}, Mass: 1e3, Radius: 10, Velocity: vect.Vect{X: 0, Y: 250}, Color: colornames.Orange},
	{Position: vect.Vect{X: 1000, Y: 0}, Mass: 1e3, Radius: 10, Velocity: vect.Vect{X: 0, Y: 180}, Color: colornames.Lightblue},
	{Position: vect.Vect{X: 1500, Y: 0}, Mass: 1e4, Radius: 25, Velocity: vect.Vect{X: 0, Y: 150}, Color: colornames.Purple},
}

func main() {
	ui := gui.NewUI("single sun", 1024, 1024)
	ui.Field.ShowTrails = true
	ui.Load(singleSun)
	ui.Field.ViewFinder.SetScale(5)
	ui.SetTime(256)
	pixelgl.Run(ui.RunGUI)
}
