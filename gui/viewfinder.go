package gui

import (
	"github.com/faiface/pixel"
	"math"
)

type ViewFinder struct {
	Offset pixel.Vec
	Scale  float64
}

func (vf *ViewFinder) SetScale(scale float64) {
	vf.Scale = math.Max(scale, 0.1)
}

func (vf ViewFinder) RealToViewFinder(input pixel.Vec) (output pixel.Vec) {
	output = input.Add(vf.Offset)
	output = output.Scaled(1 / vf.Scale)
	return
}

func (vf ViewFinder) ViewFinderToReal(input pixel.Vec) (output pixel.Vec) {
	output = input.Scaled(vf.Scale)
	output = output.Sub(vf.Offset)
	return
}
