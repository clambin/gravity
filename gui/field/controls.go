package field

import (
	"github.com/clambin/gravity/gui/object"
	"github.com/vova616/chipmunk/vect"
	"image/color"
)

// ClearObjects removes all manually added bodies from the field
func (f *Field) ClearObjects() {
	remainingObjects := make([]*object.Object, 0, len(f.objects))
	for _, o := range f.objects {
		if o.Manual {
			f.world.Space.RemoveBody(o.GetBody())
		} else {
			remainingObjects = append(remainingObjects, o)
		}
	}
	f.objects = remainingObjects
}

// ToggleShowTrails toggles the ShowTrails setting
func (f *Field) ToggleShowTrails() {
	f.ShowTrails = !f.ShowTrails
}

// Add adds a new body to the field
func (f *Field) Add(position vect.Vect, r float32, m vect.Float, velocity vect.Vect, color color.Color, manual bool) {
	o := object.New(position, r, m, velocity, f.ViewFinder, color, manual)
	f.objects = append(f.objects, o)
	f.world.Space.AddBody(o.GetBody())
}
