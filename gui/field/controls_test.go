package field_test

import (
	"github.com/clambin/gravity/gui"
	"github.com/clambin/gravity/gui/field"
	"github.com/faiface/pixel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vova616/chipmunk/vect"
	"golang.org/x/image/colornames"
	"testing"
)

func TestField_ClearObjects(t *testing.T) {
	f := field.New("test", 1000, 1000)
	f.Add(vect.Vect{X: 0, Y: 0}, 10, 1000, vect.Vect{X: 10, Y: 10}, colornames.White, false)
	f.Add(vect.Vect{X: 0, Y: 10}, 10, 1000, vect.Vect{X: 10, Y: 10}, colornames.White, true)
	stats := f.Stats()
	require.Len(t, stats, 2)

	f.ClearObjects()
	stats = f.Stats()
	require.Len(t, stats, 1)
}

func TestField_Add(t *testing.T) {
	f := field.New("test", 1000, 1000)

	f.Add(vect.Vect{X: 0, Y: 0}, 10, 1000, vect.Vect{X: 10, Y: 10}, colornames.White, true)
	stats := f.Stats()
	require.Len(t, stats, 1)
	assert.Equal(t, field.BodyStats{
		Position:     vect.Vect{X: 0, Y: 0},
		Velocity:     vect.Vect{X: 10, Y: 10},
		Acceleration: vect.Vect{X: 0, Y: 0},
	}, stats[0])
	f.ClearObjects()

	f.ViewFinder.SetScale(10)

	f.Add(vect.Vect{X: 0, Y: 0}, 10, 1000, vect.Vect{X: 100, Y: 100}, colornames.White, true)
	stats = f.Stats()
	require.Len(t, stats, 1)
	assert.Equal(t, field.BodyStats{
		Position:     vect.Vect{X: 0, Y: 0},
		Velocity:     vect.Vect{X: 100, Y: 100},
		Acceleration: vect.Vect{X: 0, Y: 0},
	}, stats[0])
	f.ClearObjects()

	f.ViewFinder.SetOffset(pixel.V(-100, -100))
	f.Add(vect.Vect{X: 0, Y: 0}, 10, 1000, vect.Vect{X: 10, Y: 10}, colornames.White, true)
	stats = f.Stats()
	require.Len(t, stats, 1)
	assert.Equal(t, field.BodyStats{
		Position:     vect.Vect{X: 0, Y: 0},
		Velocity:     vect.Vect{X: 10, Y: 10},
		Acceleration: vect.Vect{X: 0, Y: 0},
	}, stats[0])
	f.ClearObjects()

	f.ViewFinder.Reset()
	f.Add(vect.Vect{X: 0, Y: 0}, 10, 1000, vect.Vect{X: 10, Y: 10}, colornames.White, true)
	stats = f.Stats()
	require.Len(t, stats, 1)
	assert.Equal(t, field.BodyStats{
		Position:     vect.Vect{X: 0, Y: 0},
		Velocity:     vect.Vect{X: 10, Y: 10},
		Acceleration: vect.Vect{X: 0, Y: 0},
	}, stats[0])
	f.ClearObjects()
}

func TestField_Stats(t *testing.T) {
	f := field.New("test", 1000, 1000)

	for _, body := range []gui.Body{
		{Position: vect.Vect{X: 0, Y: 0}, Mass: 3.33e7, Radius: 50, Velocity: vect.Vect{X: 0, Y: 0}, Color: colornames.Yellow},
		{Position: vect.Vect{X: 200, Y: 0}, Mass: 1e2, Radius: 10, Velocity: vect.Vect{X: 0, Y: 440}, Color: colornames.Grey},
	} {
		f.Add(body.Position, body.Radius, body.Mass, body.Velocity, body.Color, false)
	}

	stats := f.Stats()
	assert.Len(t, stats, 2)

	f.Steps(100)

	stats = f.Stats()
	require.Len(t, stats, 2)
	assert.NotEqual(t, vect.Vector_Zero, stats[0].Velocity)
	assert.NotEqual(t, vect.Vector_Zero, stats[0].Acceleration)
	assert.NotEqual(t, vect.Vector_Zero, stats[1].Velocity)
	assert.NotEqual(t, vect.Vector_Zero, stats[1].Acceleration)
}

func TestField_ToggleShowTrails(t *testing.T) {
	f := field.New("test", 1000, 1000)

	assert.False(t, f.ShowTrails)
	f.ToggleShowTrails()
	assert.True(t, f.ShowTrails)
	f.ToggleShowTrails()
	assert.False(t, f.ShowTrails)
}
