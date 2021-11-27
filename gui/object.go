package gui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"gravity/particle"
)

type Object struct {
	Manual   bool
	imd      *imdraw.IMDraw
	particle *particle.Particle
	color    pixel.RGBA
	trails   []pixel.Vec
	ui       *UI
}

func NewObject(particle *particle.Particle, color pixel.RGBA, ui *UI, manual bool) (object *Object) {
	object = &Object{
		Manual:   manual,
		imd:      imdraw.New(nil),
		particle: particle,
		color:    color,
		ui:       ui,
		trails:   make([]pixel.Vec, 0),
	}
	return
}

func (object *Object) Draw(win *pixelgl.Window, showTrails bool) {
	object.imd.Clear()
	object.imd.Reset()
	object.drawObject()
	if showTrails {
		object.drawTrail()
	}
	object.drawVelocity()
	object.drawAcceleration()
	object.imd.Draw(win)
}

func (object *Object) drawObject() {
	object.imd.Color = object.color
	object.imd.Push(object.ui.Scale(pixel.V(object.particle.X, object.particle.Y)))
	object.imd.Circle(2, 0)
}

func (object *Object) drawTrail() {
	for _, trail := range object.trails {
		object.imd.Color = pixel.RGB(0, 0, 0.5)
		object.imd.Push(object.ui.Scale(pixel.V(trail.X, trail.Y)))
		object.imd.Circle(1, 0)
	}
}

func (object *Object) recordTrail() {
	object.trails = append(object.trails, pixel.V(object.particle.X, object.particle.Y))
	const maxTrails = 200
	if len(object.trails) >= maxTrails {
		object.trails = object.trails[len(object.trails)-maxTrails+1:]
	}
}

func (object *Object) drawVelocity() {
	x1 := object.particle.X
	y1 := object.particle.Y
	x2 := x1 + object.particle.VX*10
	y2 := y1 + object.particle.VY*10
	object.imd.Color = pixel.RGB(0, 0.8, 0)
	object.imd.Push(object.ui.Scale(pixel.V(x1, y1)), object.ui.Scale(pixel.V(x2, y2)))
	object.imd.Line(1)
}

func (object *Object) drawAcceleration() {
	x1 := object.particle.X
	y1 := object.particle.Y
	x2 := x1 + object.particle.AX*1000
	y2 := y1 + object.particle.AY*1000
	object.imd.Color = pixel.RGB(1, 0, 0)
	object.imd.Push(object.ui.Scale(pixel.V(x1, y1)), object.ui.Scale(pixel.V(x2, y2)))
	object.imd.Line(1)
}
