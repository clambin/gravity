package gui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk"
	"golang.org/x/image/colornames"
)

func (ui *UI) Draw(win *pixelgl.Window) {
	win.Clear(colornames.Black)
	imd := imdraw.New(nil)
	for _, body := range ui.Space.Bodies {
		if ui.showTrails {
			ui.drawTrail(imd, body)
		}
		ui.drawObject(imd, body)
		// ui.drawVelocity(imd, body)
		// ui.drawAcceleration(imd, body)
	}
	imd.Draw(win)
}

func (ui *UI) drawObject(imd *imdraw.IMDraw, body *chipmunk.Body) {
	if _, ok := ui.manualObjects[body]; ok == true {
		imd.Color = colornames.White
	} else {
		imd.Color = colornames.Purple
	}
	p := body.Position()
	imd.Push(ui.viewFinder.RealToViewFinder(pixel.V(float64(p.X), float64(p.Y))))
	imd.Circle(float64(body.Shapes[0].GetAsCircle().Radius)/ui.viewFinder.Scale, 0)
}

func (ui *UI) drawTrail(imd *imdraw.IMDraw, body *chipmunk.Body) {
	for index := range ui.trails[body] {
		if index == 0 {
			continue
		}
		imd.Color = pixel.RGB(0, 0, 0.5)
		imd.Push(ui.viewFinder.RealToViewFinder(ui.trails[body][index-1]))
		imd.Push(ui.viewFinder.RealToViewFinder(ui.trails[body][index]))
		imd.Line(1)
	}
}

func (ui *UI) drawVelocity(imd *imdraw.IMDraw, body *chipmunk.Body) {
	const magnify = 10
	p := body.Position()
	x := float64(p.X)
	y := float64(p.Y)
	v := body.Velocity()
	vx := float64(v.X * magnify)
	vy := float64(v.Y * magnify)
	imd.Color = pixel.RGB(0, 0.8, 0)
	imd.Push(ui.viewFinder.RealToViewFinder(pixel.V(x, y)), ui.viewFinder.RealToViewFinder(pixel.V(x+vx, y+vy)))
	imd.Line(1)
}

func (ui *UI) drawAcceleration(_ *imdraw.IMDraw, _ *chipmunk.Body) {
	/*	p := body.Position()
		x1 := p.X
		y1 := p.Y
		a := body.
		x2 := x1 + p.AX*1000
		y2 := y1 + p.AY*1000
		imd.Color = pixel.RGB(1, 0, 0)
		imd.Push(ui.Scale(pixel.V(x1, y1)), ui.Scale(pixel.V(x2, y2)))
		imd.Line(1)

	*/
}
