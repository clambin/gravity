package main

import (
	"github.com/faiface/pixel/pixelgl"
	"gravity/gui"
	"math/rand"
)

type particle struct {
	X  float64
	Y  float64
	M  float64
	R  float64
	VX float64
	VY float64
}

func empty() []particle {
	return []particle{}
}

func singleSun() []particle {
	return []particle{
		{X: 0, Y: 0, M: 75e8, R: 40},
		{X: 200, Y: 0, M: 10, R: 5, VY: 1.99},
		{X: 1000, Y: 0, M: 100, R: 10, VY: 1.0},
	}
}

func multipleSuns() []particle {
	return []particle{
		{X: 0, Y: 0, M: 2e10, R: 50, VX: 0, VY: 0},
		{X: -100, Y: 100, M: 2000, R: 10, VX: 0, VY: -1},
		{X: -100, Y: -100, M: 2000, R: 10, VX: 1, VY: 0},
		{X: 100, Y: -100, M: 2000, R: 10, VX: 0, VY: 1},
		{X: 100, Y: 100, M: 2000, R: 10, VX: -1, VY: 0},
	}
}

func threeBodies() []particle {
	return []particle{
		{X: 0, Y: 0, M: 2e10, R: 50, VX: 0, VY: 0},
		{X: -100, Y: 100, M: 2e10, R: 50, VX: 0, VY: -1},
		{X: -100, Y: -100, M: 2e10, R: 50, VX: 0, VY: 1},
	}
}

func ducksInARow() []particle {
	var particles []particle
	for i := -500; i <= 500; i += 100 {
		particles = append(particles, particle{
			M:  20000,
			R:  50,
			X:  float64(i),
			Y:  0,
			VX: float64(rand.Int31n(3) - 1),
			VY: float64(rand.Int31n(3) - 1),
		})
	}
	return particles
}

func main() {
	ui := gui.NewUI(1024, 1024)
	// ui.Space.Gravity = vect.Vect{X:0, Y: -0.1}

	initialState := threeBodies
	for _, body := range initialState() {
		ui.Add(body.X, body.Y, body.R, body.M, body.VX, body.VY, false)
	}

	pixelgl.Run(ui.RunGUI)
}
