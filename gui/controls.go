package gui

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"math"
)

func (ui *UI) ProcessEvents(win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyC) {
		ui.Field.ClearObjects()
	}
	if win.JustPressed(pixelgl.KeyD) {
		ui.Field.DecelerateObjects()
	}
	if win.JustPressed(pixelgl.KeyA) {
		ui.Field.AccelerateObjects()
	}
	if win.JustPressed(pixelgl.KeyT) {
		ui.Field.ToggleShowTrails()
	}
	if win.JustReleased(pixelgl.KeyRight) {
		ui.time *= 2
		fmt.Printf("time: %d\n", ui.time)
	}
	if win.JustReleased(pixelgl.KeyLeft) {
		ui.time = int(math.Max(1.0, float64(ui.time/2)))
		fmt.Printf("time: %d\n", ui.time)
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
		fmt.Printf("adding particle at (%f,%f)\n", ui.position.X, ui.position.Y)
		velocity := win.MousePosition().Sub(ui.position).Scaled(1.0 / 50)
		ui.Field.Add(ui.position, 3, 1, velocity, true)
	}
	if scroll := win.MouseScroll(); scroll.Y != 0 {
		const sensitivity = 10.0
		ui.Field.ViewFinder.SetScale(ui.Field.ViewFinder.Scale - scroll.Y/sensitivity)
		fmt.Printf("scale: %.1f\n", ui.Field.ViewFinder.Scale)
	}
}
