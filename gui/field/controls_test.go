package field_test

import (
	"github.com/faiface/pixel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vova616/chipmunk/vect"
	"gravity/gui/field"
	"testing"
)

func TestField_ClearObjects(t *testing.T) {
	f := field.New()
	f.Add(pixel.V(0, 0), 10, 1000, pixel.V(10, 10), false)
	f.Add(pixel.V(10, 10), 10, 1000, pixel.V(10, 10), true)
	stats := f.Stats()
	require.Len(t, stats, 2)

	f.ClearObjects()
	stats = f.Stats()
	require.Len(t, stats, 1)
}

func TestField_Add(t *testing.T) {
	f := field.New()

	f.Add(pixel.V(0, 0), 10, 1000, pixel.V(10, 10), true)
	stats := f.Stats()
	require.Len(t, stats, 1)
	assert.Equal(t, field.BodyStats{
		Position:     vect.Vect{X: 0, Y: 0},
		Velocity:     vect.Vect{X: 10, Y: 10},
		Acceleration: vect.Vect{X: 0, Y: 0},
	}, stats[0])
	f.ClearObjects()

	f.ViewFinder.SetScale(10)

	f.Add(pixel.V(0, 0), 10, 1000, pixel.V(10, 10), true)
	stats = f.Stats()
	require.Len(t, stats, 1)
	assert.Equal(t, field.BodyStats{
		Position:     vect.Vect{X: 0, Y: 0},
		Velocity:     vect.Vect{X: 100, Y: 100},
		Acceleration: vect.Vect{X: 0, Y: 0},
	}, stats[0])
	f.ClearObjects()

	f.ViewFinder.SetOffset(pixel.V(-100, -100))
	f.Add(pixel.V(0, 0), 10, 1000, pixel.V(10, 10), true)
	stats = f.Stats()
	require.Len(t, stats, 1)
	assert.Equal(t, field.BodyStats{
		Position:     vect.Vect{X: 1000, Y: 1000},
		Velocity:     vect.Vect{X: 100, Y: 100},
		Acceleration: vect.Vect{X: 0, Y: 0},
	}, stats[0])
	f.ClearObjects()
}

func TestField_Stats(t *testing.T) {
	f := field.New()

	f.Add(pixel.V(-10, 0), 10, 1000, pixel.V(0, 0), true)
	f.Add(pixel.V(10, 0), 10, 10, pixel.V(0, 0), true)

	stats := f.Stats()
	assert.Len(t, stats, 2)

	f.Steps(10)

	stats = f.Stats()
	require.Len(t, stats, 2)
	assert.NotEqual(t, vect.Vector_Zero, stats[0].Velocity)
	assert.NotEqual(t, vect.Vector_Zero, stats[0].Acceleration)
	assert.NotEqual(t, vect.Vector_Zero, stats[1].Velocity)
	assert.NotEqual(t, vect.Vector_Zero, stats[1].Acceleration)
}
