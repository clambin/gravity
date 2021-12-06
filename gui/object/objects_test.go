package object_test

import (
	"github.com/clambin/gravity/gui/object"
	"github.com/stretchr/testify/assert"
	"github.com/vova616/chipmunk/vect"
	"math"
	"testing"
)

func TestObject_ApplyGravity(t *testing.T) {
	objects := []*object.Object{
		object.New(vect.Vector_Zero, 100, 5e4, vect.Vector_Zero, nil, true),
		object.New(vect.Vect{X: 100, Y: 100}, 100, 1e5, vect.Vect{X: 100, Y: 100}, nil, true),
	}

	for _, o := range objects {
		o.ApplyGravity(objects)
	}

	position, velocity, acceleration := objects[0].Stats()
	assert.Equal(t, vect.Vector_Zero, position)
	assert.Equal(t, vect.Vector_Zero, velocity)
	assert.Equal(t, 4.0, math.Round(float64(acceleration.X)))
	assert.Equal(t, 4.0, math.Round(float64(acceleration.X)))

	position, velocity, acceleration = objects[1].Stats()
	assert.Equal(t, vect.Vect{X: 100, Y: 100}, position)
	assert.Equal(t, vect.Vector_Zero, velocity)
	assert.Equal(t, -2.0, math.Round(float64(acceleration.X)))
	assert.Equal(t, -2.0, math.Round(float64(acceleration.Y)))

}
