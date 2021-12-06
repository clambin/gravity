package object

import (
	"github.com/clambin/gravity/gui/viewfinder"
	"github.com/clambin/pixelmunk"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/vova616/chipmunk/vect"
	"golang.org/x/image/colornames"
)

type Object struct {
	pixelmunk.Object
	Manual     bool
	trails     []vect.Vect
	viewFinder *viewfinder.ViewFinder
}

const (
	maxTrails                 = 2000
	velocityMagnification     = 0.1
	accelerationMagnification = 0.1
)

var (
	trailColor        = pixel.RGB(0, 0, 0.5)
	objectColor       = colornames.Yellow
	manualObjectColor = colornames.White
	velocityColor     = pixel.RGB(0, 0.8, 0)
	accelerationColor = colornames.Red
)

func New(position vect.Vect, radius float32, mass vect.Float, velocity vect.Vect, viewFinder *viewfinder.ViewFinder, manual bool) (object *Object) {
	object = &Object{
		Manual:     manual,
		trails:     make([]vect.Vect, 0, maxTrails+1),
		viewFinder: viewFinder,
	}
	velocity.Sub(position)
	object.Object = pixelmunk.NewCircle(pixelmunk.ObjectOptions{
		BodyOptions: pixelmunk.ObjectBodyOptions{
			Position:   position,
			Velocity:   velocity,
			Elasticity: 0.02,
			Mass:       mass,
			CircleOptions: pixelmunk.ObjectCircleOptions{
				Radius: radius,
			},
		},
	})
	return
}

func (o Object) Stats() (position, velocity, acceleration vect.Vect) {
	body := o.Object.GetBody()
	position = body.Position()
	velocity = body.Velocity()
	acceleration = body.UserData.(vect.Vect)
	return
}

func (o *Object) RecordTrails() {
	position := o.Object.GetBody().Position()
	o.trails = append(o.trails, position)
	if len(o.trails) > maxTrails {
		o.trails = o.trails[1:]
	}
}

func (o Object) DrawTrails(imd *imdraw.IMDraw) {
	for index := range o.trails {
		if index == 0 {
			continue
		}
		imd.Color = trailColor
		imd.Push(o.viewFinder.RealToViewFinder(o.trails[index-1]))
		imd.Push(o.viewFinder.RealToViewFinder(o.trails[index]))
		imd.Line(1)
	}
}

func (o Object) DrawObject(imd *imdraw.IMDraw) {
	body := o.GetBody()
	p := body.Position()
	r := float64(body.Shapes[0].GetAsCircle().Radius) / o.viewFinder.Scale

	if o.Manual {
		imd.Color = manualObjectColor
	} else {
		imd.Color = objectColor
	}
	imd.Push(o.viewFinder.RealToViewFinder(p))
	imd.Circle(r, 0)
}

func (o Object) DrawVelocity(imd *imdraw.IMDraw) {
	body := o.GetBody()
	p := body.Position()
	v := body.Position()
	velocity := body.Velocity()
	velocity.Mult(velocityMagnification)
	v.Add(velocity)
	imd.Color = velocityColor
	imd.Push(o.viewFinder.RealToViewFinder(p), o.viewFinder.RealToViewFinder(v))
	imd.Line(1)
}

func (o Object) DrawAcceleration(imd *imdraw.IMDraw) {
	body := o.GetBody()
	p := body.Position()
	acceleration := body.UserData.(vect.Vect)
	acceleration.Mult(accelerationMagnification)
	acceleration.Add(p)
	imd.Color = accelerationColor
	imd.Push(o.viewFinder.RealToViewFinder(p), o.viewFinder.RealToViewFinder(acceleration))
	imd.Line(1)
}
