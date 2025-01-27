package gravity

import (
	"cmp"
	"github.com/jakecoffman/cp/v2"
)

// G is the gravitational constant
const G = 6.67408e-11

type Gravity struct {
	G float64
}

func (g Gravity) TotalGravitationalForce(space *cp.Space, body *cp.Body) cp.Vector {
	var force cp.Vector
	space.EachBody(func(other *cp.Body) {
		// skip body itself.  also avoids div by zero in gravitationalForce
		if other.Position() == body.Position() {
			return
		}
		force = force.Add(gravitationalForce(body, other, cmp.Or(g.G, G)))
	})
	return force
}

func gravitationalForce(body, other *cp.Body, g float64) cp.Vector {
	delta := other.Position().Sub(body.Position())
	r := delta.Length()

	// F1 = F2 = G * m1 * m2 / r^2
	f := g * (other.Mass() * body.Mass()) / (r * r)

	forceDirection := delta.Mult(1 / r)
	return forceDirection.Mult(f)
}
