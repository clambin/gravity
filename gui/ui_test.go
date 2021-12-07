package gui_test

import (
	"github.com/clambin/gravity/gui"
	"github.com/clambin/gravity/gui/field"
	"github.com/faiface/pixel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vova616/chipmunk/vect"
	"testing"
)

func TestUI_AddManual(t *testing.T) {
	ui := gui.NewUI("test", 1000, 1000)

	ui.AddManual(pixel.V(0, 0), pixel.V(10, 10), 10, 1000)
	stats := ui.Field.Stats()
	require.Len(t, stats, 1)
	assert.Equal(t, field.BodyStats{
		Position:     vect.Vect{X: 0, Y: 0},
		Velocity:     vect.Vect{X: 10, Y: 10},
		Acceleration: vect.Vect{X: 0, Y: 0},
	}, stats[0])
	ui.Field.ClearObjects()

	ui.Field.ViewFinder.SetScale(10)

	ui.AddManual(pixel.V(0, 0), pixel.V(10, 10), 10, 1000)
	stats = ui.Field.Stats()
	require.Len(t, stats, 1)
	assert.Equal(t, field.BodyStats{
		Position:     vect.Vect{X: 0, Y: 0},
		Velocity:     vect.Vect{X: 100, Y: 100},
		Acceleration: vect.Vect{X: 0, Y: 0},
	}, stats[0])
	ui.Field.ClearObjects()

	ui.Field.ViewFinder.SetOffset(pixel.V(-100, -100))
	ui.AddManual(pixel.V(0, 0), pixel.V(10, 10), 10, 1000)
	stats = ui.Field.Stats()
	require.Len(t, stats, 1)
	assert.Equal(t, field.BodyStats{
		Position:     vect.Vect{X: 1000, Y: 1000},
		Velocity:     vect.Vect{X: 100, Y: 100},
		Acceleration: vect.Vect{X: 0, Y: 0},
	}, stats[0])
	ui.Field.ClearObjects()

	ui.Field.ViewFinder.Reset()
	ui.AddManual(pixel.V(0, 0), pixel.V(10, 10), 10, 1000)
	stats = ui.Field.Stats()
	require.Len(t, stats, 1)
	assert.Equal(t, field.BodyStats{
		Position:     vect.Vect{X: 0, Y: 0},
		Velocity:     vect.Vect{X: 10, Y: 10},
		Acceleration: vect.Vect{X: 0, Y: 0},
	}, stats[0])
	ui.Field.ClearObjects()
}

func TestUI_Load(t *testing.T) {
	ui := gui.NewUI("test", 1000, 1000)

	ui.Load([]gui.Body{
		{Position: vect.Vect{X: 0, Y: 0}, Mass: 15e6, Radius: 100},
		{Position: vect.Vect{X: -2500, Y: 2500}, Mass: 15e5, Radius: 50, Velocity: vect.Vect{Y: -25}},
		{Position: vect.Vect{X: 2500, Y: -2500}, Mass: 15e5, Radius: 50, Velocity: vect.Vect{Y: 25}},
	})

	stats := ui.Field.Stats()
	assert.Equal(t, []field.BodyStats{
		{Position: vect.Vect{X: 0, Y: 0}, Velocity: vect.Vect{X: 0, Y: 0}, Acceleration: vect.Vect{X: 0, Y: 0}},
		{Position: vect.Vect{X: -2500, Y: 2500}, Velocity: vect.Vect{X: 0, Y: -25}, Acceleration: vect.Vect{X: 0, Y: 0}},
		{Position: vect.Vect{X: 2500, Y: -2500}, Velocity: vect.Vect{X: 0, Y: 25}, Acceleration: vect.Vect{X: 0, Y: 0}},
	}, stats)
}
