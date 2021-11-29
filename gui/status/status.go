package status

import (
	"fmt"
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
	CanvasWidthOffset  = 7
	CanvasHeight       = 21
	CanvasHeightOffset = 14
)

// MakeCanvas creates the pixelgl canvas on which to draw the stats
func (reporter *Reporter) MakeCanvas(width, height float64) {
	reporter.canvas = pixelgl.NewCanvas(pixel.Rect{
		Min: pixel.Vec{X: -CanvasWidthOffset, Y: -CanvasHeight + CanvasHeightOffset},
		Max: pixel.Vec{X: width - CanvasWidthOffset, Y: CanvasHeightOffset},
	})
	reporter.offset = pixel.V(0, float64(-height+CanvasHeight)/2)
}

// Draw draws the stats canvas and adds it to the Target
func (reporter Reporter) Draw(win pixel.Target, offset pixel.Vec, scale float64, speed int) {
	reporter.canvas.Clear(colornames.Blue)
	reporter.writeStats(reporter.canvas, offset, scale, speed)
	reporter.canvas.Draw(win, pixel.IM.Moved(reporter.offset))
}

func (reporter Reporter) writeStats(win pixel.Target, offset pixel.Vec, scale float64, speed int) {
	t := text.New(pixel.V(0, 0), text.Atlas7x13)
	_, _ = fmt.Fprintf(t, "offset: (%.0f,%.0f) scale: %.1f speed: %d", offset.X, offset.Y, scale, speed)
	t.Draw(win, pixel.IM)
}
