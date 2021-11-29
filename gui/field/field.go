package field

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
)

// Field represents a gravity simulation field
type Field struct {
	ShowTrails    bool
	ViewFinder    ViewFinder
	space         *chipmunk.Space
	manualObjects map[*chipmunk.Body]struct{}
	trails        map[*chipmunk.Body][]pixel.Vec
	canvas        *pixelgl.Canvas
}

// New creates a gravity simulation field
func New() *Field {
	return &Field{
		space:         chipmunk.NewSpace(),
		manualObjects: make(map[*chipmunk.Body]struct{}),
		trails:        make(map[*chipmunk.Body][]pixel.Vec),
		ViewFinder:    ViewFinder{Scale: 1},
	}
}

// Steps performs a number of steps of the simulation
func (f *Field) Steps(n int) {
	for i := 0; i < n; i++ {
		f.Step()
	}
	f.RecordTrails()
}

// Step performs one step of the simulation
func (f *Field) Step() {
	f.space.Step(1.0)
	f.gravity()
}

// RecordTrails records the position of each body, so a trail can be drawn
func (f *Field) RecordTrails() {
	const maxTrails = 2000
	for _, body := range f.space.Bodies {
		trails, _ := f.trails[body]
		p := body.Position()
		trails = append(trails, pixel.V(float64(p.X), float64(p.Y)))
		if len(trails) > maxTrails {
			trails = trails[1:]
		}
		f.trails[body] = trails
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
	for _, body := range f.space.Bodies {
		var acceleration vect.Vect
		if body.UserData != nil {
			acceleration = body.UserData.(vect.Vect)
		}
		output = append(output, BodyStats{
			Position:     body.Position(),
			Velocity:     body.Velocity(),
			Acceleration: acceleration,
		})
	}
	return
}
