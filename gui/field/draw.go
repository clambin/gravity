package field

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

// MakeCanvas creates a pixelgl canvas on which to draw the field
func (f *Field) MakeCanvas() {
	f.canvas = pixelgl.NewCanvas(f.world.Bounds)
}

// Draw draws the bodies on the specified Target
func (f Field) Draw(win pixel.Target) {
	f.canvas.Clear(pixel.RGB(0.01, 0.01, 0.01))
	imd := imdraw.New(nil)
	f.drawTrails(imd)
	f.drawObjects(imd)
	f.drawVelocity(imd)
	f.drawAcceleration(imd)
	imd.Draw(f.canvas)
	f.canvas.Draw(win, pixel.IM)
}

func (f Field) drawTrails(imd *imdraw.IMDraw) {
	if f.ShowTrails {
		for _, o := range f.objects {
			o.DrawTrails(imd)
		}
	}
}

func (f Field) drawObjects(imd *imdraw.IMDraw) {
	for _, o := range f.objects {
		o.DrawObject(imd)
	}
}

func (f Field) drawVelocity(imd *imdraw.IMDraw) {
	for _, o := range f.objects {
		o.DrawVelocity(imd)
	}
}

func (f Field) drawAcceleration(imd *imdraw.IMDraw) {
	for _, o := range f.objects {
		o.DrawAcceleration(imd)
	}
}
