package main

import (
	"github.com/faiface/pixel/pixelgl"
	"gravity/gui"
	"gravity/particle"
	"math/rand"
)

func empty() []*particle.Particle {
	return []*particle.Particle{}
}

func singleSun() []*particle.Particle {
	return []*particle.Particle{
		{X: 0, Y: 0, M: 75000},
		{X: 100, Y: 0, M: 10, VY: 0.9},
		{X: 1000, Y: 0, M: 100, VY: 0.3},
	}
}

func multipleSuns() []*particle.Particle {
	return []*particle.Particle{
		{X: 0, Y: 0, M: 200000, VX: 0, VY: 0},
		{X: -100, Y: 100, M: 200, VX: 0, VY: -1},
		{X: -100, Y: -100, M: 200, VX: 1, VY: 0},
		{X: 100, Y: -100, M: 200, VX: 0, VY: 1},
		{X: 100, Y: 100, M: 200, VX: -1, VY: 0},
	}
}

func ducksInARow() []*particle.Particle {
	var particles []*particle.Particle
	for i := -500; i <= 500; i += 100 {
		particles = append(particles, &particle.Particle{
			M:  20000,
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

	initialState := singleSun
	for _, body := range initialState() {
		ui.Add(body, false)
	}

	pixelgl.Run(ui.RunGUI)
}
