package field

import (
	"github.com/faiface/pixel"
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
)

// ClearObjects removes all manually added bodies from the field
func (f *Field) ClearObjects() {
	var bodies []*chipmunk.Body
	for _, body := range f.space.Bodies {
		if _, ok := f.manualObjects[body]; ok == false {
			bodies = append(bodies, body)
		}
	}
	f.space.Bodies = bodies
	f.manualObjects = make(map[*chipmunk.Body]struct{})
}

// ToggleShowTrails toggles the ShowTrails setting
func (f *Field) ToggleShowTrails() {
	f.ShowTrails = !f.ShowTrails
}

/*
// DecelerateBodies reduces the velocity of each body
func (f *Field) DecelerateBodies() {
	f.alterVelocity(0.9)
}

// AccelerateBodies increases the velocity of each body
func (f *Field) AccelerateBodies() {
	f.alterVelocity(1.1)
}

func (f *Field) alterVelocity(factor vect.Float) {
	for _, body := range f.space.Bodies {
		v := body.Velocity()
		v.Mult(factor)
		body.SetVelocity(float32(v.X), float32(v.Y))
	}
}
*/

// Add adds a new body to the field
func (f *Field) Add(position pixel.Vec, r, m float32, velocity pixel.Vec, manual bool) {
	// convert position & velocity to real coordinates
	p := f.ViewFinder.ViewFinderToReal(position)
	v := velocity.Scaled(f.ViewFinder.Scale)

	circle := chipmunk.NewCircle(vect.Vector_Zero, r)
	//circle.SetElasticity(0.1)
	body := chipmunk.NewBody(vect.Float(m), circle.Moment(m))
	body.SetPosition(vect.Vect{X: vect.Float(p.X), Y: vect.Float(p.Y)})
	body.AddShape(circle)
	body.SetVelocity(float32(v.X), float32(v.Y))
	f.space.AddBody(body)

	if manual {
		f.manualObjects[body] = struct{}{}
	}
}
