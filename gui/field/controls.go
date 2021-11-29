package field

import (
	"fmt"
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

func (f *Field) DecelerateObjects() {
	/*
		for _, p := range ui.Space.Particles {
			//if object.Manual == true {
			p.AX /= 2
			p.AY /= 2
			fmt.Printf("decelerate: (%f,%f)\n", p.AX, p.AY)
			//}
		}
	*/
}

func (f *Field) AccelerateObjects() {
	/*
		for _, p := range ui.Space.Particles {
			// if object.Manual == true {
			p.AX *= 2
			p.AY *= 2
			fmt.Printf("accelerate: (%f,%f)\n", p.AX, p.AY)
			//}
		}
	*/
}

// Add adds a new body to the field
func (f *Field) Add(position pixel.Vec, r, m float32, velocity pixel.Vec, manual bool) {
	// convert position & velocity to real coordinates
	p := f.ViewFinder.ViewFinderToReal(position)
	v := velocity.Scaled(f.ViewFinder.Scale)

	fmt.Printf("viewfinder: offset: (%.1f/%.1f), scale: %f\n",
		f.ViewFinder.Offset.X, f.ViewFinder.Offset.Y,
		f.ViewFinder.Scale,
	)
	fmt.Printf("position: (%1.f/%.1f) / (%.1f/%.1f), velocity: (%.1f/%.1f) / (%.1f/%.1f)\n",
		position.X, position.Y,
		p.X, p.Y,
		v.X, v.Y,
		v.X, v.Y,
	)

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
