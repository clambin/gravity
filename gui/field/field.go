package field

import (
	"github.com/clambin/gravity/gui/object"
	"github.com/clambin/gravity/gui/viewfinder"
	"github.com/clambin/pixelmunk"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk/vect"
)

// Field represents a gravity simulation field
type Field struct {
	ShowTrails bool
	ViewFinder *viewfinder.ViewFinder
	objects    []*object.Object
	world      *pixelmunk.World
	canvas     *pixelgl.Canvas
}

// New creates a gravity simulation field
func New(name string, x, y float64) *Field {
	return &Field{
		ViewFinder: &viewfinder.ViewFinder{Scale: 1},
		objects:    make([]*object.Object, 0),
		world:      pixelmunk.NewWorld(name, -x/2, -y/2, x/2, y/2),
	}
}

// Steps performs a number of steps of the simulation
func (f *Field) Steps(n int) {
	for i := 0; i < n; i++ {
		f.Step()
	}
	f.recordTrails()
}

// Step performs one step of the simulation
func (f *Field) Step() {
	f.world.Space.Step(0.001)
	f.gravity()
}

// recordTrails records the position of each body, so a trail can be drawn
func (f *Field) recordTrails() {
	for _, o := range f.objects {
		o.RecordTrails()
	}
}

// gravity applies the gravitational force from all other objects on each object
func (f *Field) gravity() {
	for _, o := range f.objects {
		o.ApplyGravity(f.objects)
	}
}

// BodyStats holds statistics for a body on the field
type BodyStats struct {
	Position     vect.Vect
	Velocity     vect.Vect
	Acceleration vect.Vect
}

// Stats generates BodyStats for each body on the field
func (f Field) Stats() (output []BodyStats) {
	for _, o := range f.objects {
		body := o.GetBody()
		var acceleration vect.Vect
		if body.UserData != nil {
			acceleration = o.GetBody().UserData.(vect.Vect)
		}
		output = append(output, BodyStats{
			Position:     body.Position(),
			Velocity:     body.Velocity(),
			Acceleration: acceleration,
		})
	}
	return
}
