package stats

import (
	"fmt"
	"github.com/clambin/gravity/gui/field"
	"github.com/clambin/gravity/gui/status"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

type Reporter struct {
	canvas *pixelgl.Canvas
	offset pixel.Vec
}

const (
	CanvasWidth        = 400
	CanvasWidthOffset  = 7
	CanvasHeight       = 200
	CanvasHeightOffset = 14
)

// MakeCanvas creates the pixelgl canvas on which to draw the stats
func (reporter *Reporter) MakeCanvas(x, y float64) {
	reporter.canvas = pixelgl.NewCanvas(pixel.Rect{
		Min: pixel.Vec{X: -CanvasWidthOffset, Y: -CanvasHeight + CanvasHeightOffset},
		Max: pixel.Vec{X: CanvasWidth - CanvasWidthOffset, Y: CanvasHeightOffset},
	})
	reporter.offset = pixel.V(float64(x-CanvasWidth)/2, float64(-y+CanvasHeight)/2+status.CanvasHeight)
}

// Draw draws the stats canvas and adds it to the Target
func (reporter Reporter) Draw(win pixel.Target, stats []field.BodyStats) {
	reporter.canvas.Clear(colornames.Darkslategrey)
	reporter.writeStats(reporter.canvas, stats)
	reporter.canvas.Draw(win, pixel.IM.Moved(reporter.offset))
}

func (reporter Reporter) writeStats(win pixel.Target, stats []field.BodyStats) {
	t := text.New(pixel.V(0, 0), text.Atlas7x13)
	for index, body := range stats {
		p := body.Position
		v := body.Velocity.Length()
		a := body.Acceleration.Length()
		_, _ = fmt.Fprintf(t, "#%2d - v: %4.1f - a:%4.0f - p:(%+.0f,%+.0f)\n",
			index+1, v, 1000*a, p.X, p.Y)
	}
	t.Draw(win, pixel.IM)
}
