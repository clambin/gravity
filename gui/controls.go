package gui

import (
	"github.com/faiface/pixel/pixelgl"
	"math"
)

func (ui *UI) ProcessEvents(win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyC) {
		ui.Field.ClearObjects()
	}
	if win.JustPressed(pixelgl.KeyT) {
		ui.Field.ToggleShowTrails()
	}
	if win.JustReleased(pixelgl.KeyRight) {
		ui.time *= 2
	}
	if win.JustReleased(pixelgl.KeyLeft) {
		ui.time = int(math.Max(1.0, float64(ui.time/2)))
	}
	if win.JustReleased(pixelgl.Key0) {
		ui.Field.ViewFinder.Reset()
	}
	if win.JustPressed(pixelgl.MouseButtonLeft) {
		ui.position = win.MousePosition()
	} else if win.Pressed(pixelgl.MouseButtonLeft) && !win.Pressed(pixelgl.KeyLeftControl) {
		newPosition := win.MousePosition()
		ui.Field.ViewFinder.SetOffset(newPosition.Sub(ui.position))
		ui.position = newPosition
	} else if win.JustReleased(pixelgl.MouseButtonLeft) && win.Pressed(pixelgl.KeyLeftControl) {
		ui.AddManual(ui.position, win.MousePosition(), 5, 1)
	}
	if scroll := win.MouseScroll(); scroll.Y != 0 {
		const sensitivity = 10.0
		ui.Field.ViewFinder.SetScale(ui.Field.ViewFinder.Scale - scroll.Y/sensitivity)
	}
}
