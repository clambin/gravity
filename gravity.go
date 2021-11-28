package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"gravity/gui"
	"math/rand"
)

type particle struct {
	X  float64
	Y  float64
	M  float32
	R  float32
	VX float64
	VY float64
}

func empty() []particle {
	return []particle{}
}

func singleSun() []particle {
	return []particle{
		{X: 0, Y: 0, M: 75e8, R: 40},
		{X: 500, Y: 0, M: 100, R: 10, VX: -0.0, VY: 1.6},
		{X: 1000, Y: 0, M: 100, R: 10, VY: 1.0},
	}
}

func multipleSuns() []particle {
	return []particle{
		{X: 0, Y: 0, M: 2e10, R: 100, VX: 0, VY: 0},
		{X: -100, Y: 100, M: 2e7, R: 30, VX: 0, VY: -1},
		{X: -100, Y: -100, M: 2e7, R: 30, VX: 1, VY: 0},
		{X: 100, Y: -100, M: 2e7, R: 30, VX: 0, VY: 1},
		{X: 100, Y: 100, M: 2e7, R: 30, VX: -1, VY: 0},
	}
}

func threeBodies() []particle {
	return []particle{
		{X: 0, Y: 0, M: 2e10, R: 50, VX: 0, VY: 0},
		{X: -1000, Y: 1000, M: 2e10, R: 50, VX: 0, VY: -1},
		{X: -1000, Y: -1000, M: 2e10, R: 50, VX: 0, VY: 1},
	}
}

func chaos() []particle {
	var particles []particle

	particles = append(particles, particle{
		X:  -1000,
		Y:  -1000,
		M:  1e12,
		R:  30,
		VX: 10,
		VY: 10,
	})
	for i := 0; i < 100; i++ {
		particles = append(particles, particle{
			X:  float64(rand.Int31n(2000) - 1000),
			Y:  float64(rand.Int31n(2000) - 1000),
			M:  1e3,
			R:  15,
			VX: float64(rand.Int31n(5) - 2),
			VY: float64(rand.Int31n(5) - 2),
		})
	}
	return particles
}

func main() {
	ui := gui.NewUI(1024, 1024)
	initialState := singleSun()
	for _, body := range initialState {
		ui.Field.Add(pixel.V(body.X, body.Y), body.R, body.M, pixel.V(body.VX, body.VY), false)
	}
	pixelgl.Run(ui.RunGUI)
}
