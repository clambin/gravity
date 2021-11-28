package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk/vect"
	"gravity/gui"
	"math/rand"
)

type particle struct {
	X  vect.Float
	Y  vect.Float
	M  float32
	R  float32
	VX float32
	VY float32
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
		{X: -1000, Y: 1000, M: 2e10, R: 50, VX: 0, VY: -1},
		{X: -1000, Y: -1000, M: 2e10, R: 50, VX: 0, VY: 1},
	}
}

func ducksInARow() []particle {
	var particles []particle
	for i := -500; i <= 500; i += 100 {
		particles = append(particles, particle{
			M:  20000,
			R:  50,
			X:  vect.Float(i),
			Y:  0,
			VX: float32(rand.Int31n(3) - 1),
			VY: float32(rand.Int31n(3) - 1),
		})
	}
	return particles
}

func main() {
	ui := gui.NewUI(1024, 1024)
	// ui.Space.Gravity = vect.Vect{X:0, Y: -0.1}

	initialState := threeBodies
	for _, body := range initialState() {
		ui.Add(vect.Vect{X: body.X, Y: body.Y}, body.R, body.M, body.VX, body.VY, false)
	}

	pixelgl.Run(ui.RunGUI)
}
