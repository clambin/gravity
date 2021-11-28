package field

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
	"golang.org/x/image/colornames"
)

// MakeCanvas creates a pixelgl canvas on which to draw the field
func (f *Field) MakeCanvas(x, y float64) {
	f.canvas = pixelgl.NewCanvas(pixel.Rect{Min: pixel.Vec{X: -x / 2, Y: -y / 2}, Max: pixel.Vec{X: x / 2, Y: y / 2}})
}

// Draw draws the bodies on the specified Target
func (f Field) Draw(win pixel.Target) {
	f.canvas.Clear(pixel.RGB(0.01, 0.01, 0.01))
	imd := imdraw.New(nil)
	f.drawSpace(imd)
	imd.Draw(f.canvas)
	f.canvas.Draw(win, pixel.IM)
}

func (f Field) drawSpace(imd *imdraw.IMDraw) {
	for _, body := range f.space.Bodies {
		if f.ShowTrails {
			f.drawTrail(imd, body)
		}
	}
	for _, body := range f.space.Bodies {
		f.drawObject(imd, body)
		f.drawVelocity(imd, body)
		f.drawAcceleration(imd, body)
	}
}

func (f Field) drawObject(imd *imdraw.IMDraw, body *chipmunk.Body) {
	if _, ok := f.manualObjects[body]; ok == true {
		imd.Color = colornames.White
	} else {
		imd.Color = colornames.Yellow
	}
	p := body.Position()
	imd.Push(f.ViewFinder.RealToViewFinder(pixel.V(float64(p.X), float64(p.Y))))
	imd.Circle(float64(body.Shapes[0].GetAsCircle().Radius)/f.ViewFinder.Scale, 0)
}

func (f Field) drawTrail(imd *imdraw.IMDraw, body *chipmunk.Body) {
	for index := range f.trails[body] {
		if index == 0 {
			continue
		}
		imd.Color = pixel.RGB(0, 0, 0.5)
		imd.Push(f.ViewFinder.RealToViewFinder(f.trails[body][index-1]))
		imd.Push(f.ViewFinder.RealToViewFinder(f.trails[body][index]))
		imd.Line(1)
	}
}

func (f Field) drawVelocity(imd *imdraw.IMDraw, body *chipmunk.Body) {
	const magnification = 10
	p := body.Position()
	x := float64(p.X)
	y := float64(p.Y)
	v := body.Velocity()
	vx := float64(v.X * magnification)
	vy := float64(v.Y * magnification)
	imd.Color = pixel.RGB(0, 0.8, 0)
	imd.Push(f.ViewFinder.RealToViewFinder(pixel.V(x, y)), f.ViewFinder.RealToViewFinder(pixel.V(x+vx, y+vy)))
	imd.Line(1)
}

func (f Field) drawAcceleration(imd *imdraw.IMDraw, body *chipmunk.Body) {
	p := body.Position()
	force := body.UserData.(vect.Vect)
	x1 := float64(p.X)
	y1 := float64(p.Y)
	const magnification = 10000
	x2 := x1 + magnification*float64(force.X/body.Mass())
	y2 := y1 + magnification*float64(force.Y/body.Mass())
	imd.Color = pixel.RGB(1, 0, 0)
	imd.Push(f.ViewFinder.RealToViewFinder(pixel.V(x1, y1)), f.ViewFinder.RealToViewFinder(pixel.V(x2, y2)))
	imd.Line(1)
}
