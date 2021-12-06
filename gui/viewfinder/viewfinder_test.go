package viewfinder_test

import (
	"github.com/clambin/gravity/gui/viewfinder"
	"github.com/faiface/pixel"
	"github.com/stretchr/testify/assert"
	"github.com/vova616/chipmunk/vect"
	"testing"
)

func TestViewFinder_RealToViewFinder(t *testing.T) {
	vf := &viewfinder.ViewFinder{Scale: 1}

	input := vect.Vect{X: 1000, Y: 500}
	output := vf.RealToViewFinder(input)
	assert.Equal(t, pixel.V(1000, 500), output)

	vf.SetScale(100)
	output = vf.RealToViewFinder(input)
	assert.Equal(t, pixel.V(10, 5), output)

	vf.Offset = pixel.V(-100, 100)
	output = vf.RealToViewFinder(input)
	assert.Equal(t, pixel.V(9, 6), output)
}

func TestViewFinder_ViewFinderToReal(t *testing.T) {
	vf := &viewfinder.ViewFinder{Scale: 1}

	input := pixel.V(10, 5)
	output := vf.ViewFinderToReal(input)
	assert.Equal(t, vect.Vect{X: 10, Y: 5}, output)

	vf.SetScale(100)
	output = vf.ViewFinderToReal(input)
	assert.Equal(t, vect.Vect{X: 1000, Y: 500}, output)

	vf.Offset = pixel.V(-100, 100)
	output = vf.ViewFinderToReal(input)
	assert.Equal(t, vect.Vect{X: 1100, Y: 400}, output)
}
