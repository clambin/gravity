package viewfinder

import (
	"github.com/faiface/pixel"
	"github.com/vova616/chipmunk/vect"
	"math"
)

// ViewFinder allows the user to zoom or move the view on the field
type ViewFinder struct {
	Offset pixel.Vec
	Scale  float64
}

// SetScale zooms in or out of the field
func (vf *ViewFinder) SetScale(scale float64) {
	vf.Scale = math.Max(scale, 0.1)
}

// SetOffset sets an offset for viewing the field
func (vf *ViewFinder) SetOffset(offset pixel.Vec) {
	vf.Offset = vf.Offset.Add(offset.Scaled(vf.Scale))
}

// Reset resets the viewfinder settings
func (vf *ViewFinder) Reset() {
	vf.Offset = pixel.V(0, 0)
	vf.Scale = 1
}

// RealToViewFinder converts a set of real coordinates to coordinates on the screen
func (vf ViewFinder) RealToViewFinder(input vect.Vect) (output pixel.Vec) {
	output = pixel.V(float64(input.X), float64(input.Y))
	output = output.Add(vf.Offset)
	output = output.Scaled(1 / vf.Scale)
	return
}

// ViewFinderToReal converts a set of coordinates on the screen to real coordinates
func (vf ViewFinder) ViewFinderToReal(input pixel.Vec) (output vect.Vect) {
	output = vect.Vect{X: vect.Float(input.X), Y: vect.Float(input.Y)}
	output.Mult(vect.Float(vf.Scale))
	output.Sub(vect.Vect{X: vect.Float(vf.Offset.X), Y: vect.Float(vf.Offset.Y)})
	return
}
