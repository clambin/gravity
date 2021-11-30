package main

import (
	"github.com/clambin/gravity/gui"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math/rand"
)

// callback will keep the viewfinder focused on the first body
func callback(ui *gui.UI) {
	b := ui.Field.Stats()
	if len(b) > 0 {
		central := b[0].Position
		v := pixel.V(-float64(central.X), -float64(central.Y))
		ui.Field.ViewFinder.Offset = v
	}
}

func bodies() (result []gui.Body) {
	for i := 0; i < 1; i++ {
		result = append(result, gui.Body{
			X: float64(rand.Int31n(1000) - 500),
			Y: float64(rand.Int31n(1000) - 500),
			R: 100,
			M: 1e10,
		})
	}

	for i := 0; i < 10; i++ {
		result = append(result, gui.Body{
			X: float64(rand.Int31n(1000) - 500),
			Y: float64(rand.Int31n(1000) - 500),
			R: 50,
			M: 1e7,
		})
	}

	for i := 0; i < 100; i++ {
		result = append(result, gui.Body{
			X: float64(rand.Int31n(1000) - 500),
			Y: float64(rand.Int31n(1000) - 500),
			R: 10,
			M: 1e3,
		})
	}

	return
}

func main() {
	ui := gui.NewUI(1024, 1024)
	ui.Load(bodies())
	ui.Callback = callback
	ui.Field.ViewFinder.SetScale(30)
	ui.Field.ShowTrails = true
	ui.SetTime(512)
	pixelgl.Run(ui.RunGUI)
}
