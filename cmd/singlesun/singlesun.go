package main

import (
	"github.com/clambin/gravity/gui"
	"github.com/faiface/pixel/pixelgl"
)

var singleSun = []gui.Body{
	{X: 0, Y: 0, M: 75e8, R: 40},
	{X: 500, Y: 0, M: 1e4, R: 10, VX: -0.0, VY: 1.4},
	{X: 1000, Y: 0, M: 1e4, R: 10, VY: 1.0},
	{X: 5000, Y: 0, M: 5e7, R: 25, VY: 0.5},
}

func main() {
	ui := gui.NewUI(1024, 1024)
	ui.Load(singleSun)
	pixelgl.Run(ui.RunGUI)
}
