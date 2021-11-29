package field

import (
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
)

const G = 0.0000001

func (f *Field) gravity() {
	for _, body := range f.space.Bodies {
		f.applyGravity(body)
	}
}

func (f *Field) applyGravity(body *chipmunk.Body) {
	var fx, fy vect.Float
	for _, other := range f.space.Bodies {
		if other == body {
			continue
		}

		dx := other.Position().X - body.Position().X
		dy := other.Position().Y - body.Position().Y
		r := vect.Vect{X: dx, Y: dy}.Length()
		// F = m * a
		// => a = F / m
		// F1 = F2 = G * m1 * m2 / r^2
		// => a1 = G * m2 / r2
		//
		// ax = dx * a / r
		// => ax = dx * G * m2 / r^3
		r3 := r * r * r
		fx += dx * G * (other.Mass() * body.Mass()) / r3
		fy += dy * G * (other.Mass() * body.Mass()) / r3
	}
	body.AddForce(float32(fx), float32(fy))
	acceleration := vect.Vect{X: fx, Y: fy}
	acceleration.Mult(1 / body.Mass())
	body.UserData = acceleration
}
