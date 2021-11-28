package stats

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"gravity/gui/field"
)

type Stats struct {
	canvas *pixelgl.Canvas
	offset pixel.Vec
}

const (
	statsX       = 300
	statsY       = 200
	statsYOffset = 11
)

// MakeCanvas creates the pixelgl canvas on which to draw the stats
func (s *Stats) MakeCanvas(x, y float64) {
	s.canvas = pixelgl.NewCanvas(pixel.Rect{Min: pixel.Vec{X: 0, Y: -statsY + statsYOffset}, Max: pixel.Vec{X: statsX, Y: statsYOffset}})
	s.offset = pixel.V(float64(x-statsX)/2, float64(-y+statsY)/2)
}

// Draw draws the stats canvas and adds it to the Target
func (s Stats) Draw(win pixel.Target, stats []field.BodyStats) {
	s.canvas.Clear(colornames.Darkslategrey)
	s.writeStats(s.canvas, stats)
	s.canvas.Draw(win, pixel.IM.Moved(s.offset))
}

func (s Stats) writeStats(win pixel.Target, stats []field.BodyStats) {
	t := text.New(pixel.V(0, 0), text.Atlas7x13)
	for index, body := range stats {
		p := body.Position
		v := body.Velocity.Length()
		a := body.Acceleration
		_, _ = fmt.Fprintf(t, "#%2d - v: %4.1f - p:(%.0f,%.0f) - a:(%.0f,%.0f)\n", index+1, v, p.X, p.Y, a.X, a.Y)
	}
	t.Draw(win, pixel.IM)
}
