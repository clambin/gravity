package field

import (
	"github.com/faiface/pixel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestField_Steps(t *testing.T) {
	f := New()
	f.MaxTrails = 100

	f.Add(pixel.V(0, 0), 10, 7e8, pixel.V(0, 0), false)
	f.Add(pixel.V(100, 0), 10, 1e3, pixel.V(0.5, 0), false)

	for i := 0; i < 110; i++ {
		f.Steps(1)
	}

	require.Len(t, f.space.Bodies, 2)
	require.Len(t, f.trails, 2)
	for _, trails := range f.trails {
		assert.Len(t, trails, 100)
	}
}
