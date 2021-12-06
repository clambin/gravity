package object

import "github.com/vova616/chipmunk/vect"

//const G = 6.67408e-11
const G = 1

func (o *Object) ApplyGravity(objects []*Object) {
	var force vect.Vect
	for _, other := range objects {
		if o != other {
			force.Add(o.calculateGravity(other))
		}
	}
	o.GetBody().AddForce(float32(force.X), float32(force.Y))

	acceleration := force
	acceleration.Mult(1 / o.GetBody().Mass())
	o.GetBody().UserData = acceleration
}

func (o *Object) calculateGravity(other *Object) (force vect.Vect) {
	dx := other.GetBody().Position().X - o.GetBody().Position().X
	dy := other.GetBody().Position().Y - o.GetBody().Position().Y
	r := vect.Vect{X: dx, Y: dy}.Length()
	// F1 = F2 = G * m1 * m2 / r^2
	f := G * (other.GetBody().Mass() * o.GetBody().Mass()) / (r * r)
	//
	// fx = dx * f / r
	// fy = dy * f / r
	force = vect.Vect{X: dx, Y: dy}
	force.Mult(f / r)
	// force.X = dx * f / r
	// force.Y = dy * f / r

	return
}
