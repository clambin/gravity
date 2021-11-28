package gui

import (
	"github.com/vova616/chipmunk"
	"math"
)

func (ui *UI) gravity() {
	for _, body := range ui.Space.Bodies {
		ui.applyGravity(body)
	}
}

func distance(a, b *chipmunk.Body) float32 {
	dx := float64(a.Position().X - b.Position().X)
	dy := float64(a.Position().Y - b.Position().Y)
	return float32(math.Sqrt(dx*dx + dy*dy))
}

func (ui *UI) applyGravity(body *chipmunk.Body) {
	var fx, fy float32
	for _, other := range ui.Space.Bodies {
		if other == body {
			continue
		}

		dx := float32(other.Position().X - body.Position().X)
		dy := float32(other.Position().Y - body.Position().Y)
		r := distance(body, other)
		// F = m * a
		// => a = F / m
		// F1 = F2 = G * m1 * m2 / r^2
		// => a1 = G * m2 / r2
		//
		// ax = dx * a / r
		// => ax = dx * G * m2 / r^3
		r3 := r * r * r
		const G = 0.0000001
		fx += dx * G * (float32(other.Mass() * body.Mass())) / r3
		fy += dy * G * (float32(other.Mass() * body.Mass())) / r3
	}
	body.AddForce(fx, fy)
}
