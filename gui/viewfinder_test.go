package gui_test

import (
	"github.com/faiface/pixel"
	"github.com/stretchr/testify/assert"
	"gravity/gui"
	"testing"
)

func TestViewFinder_RealToViewFinder(t *testing.T) {
	vf := &gui.ViewFinder{Scale: 1}

	input := pixel.V(1000, 500)
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
	vf := &gui.ViewFinder{Scale: 1}

	input := pixel.V(10, 5)
	output := vf.ViewFinderToReal(input)
	assert.Equal(t, pixel.V(10, 5), output)

	vf.SetScale(100)
	output = vf.ViewFinderToReal(input)
	assert.Equal(t, pixel.V(1000, 500), output)

	vf.Offset = pixel.V(-100, 100)
	output = vf.ViewFinderToReal(input)
	assert.Equal(t, pixel.V(1100, 400), output)
}
