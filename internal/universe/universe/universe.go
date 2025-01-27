package universe

import (
	"fmt"
	"github.com/clambin/gravity/gravity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/jakecoffman/cp/v2"
	"golang.org/x/image/colornames"
	"image/color"
	"strings"
)

type Universe struct {
	FocusObject *cp.Shape
	space       *cp.Space
	gravity.Gravity
	zoom         float64
	speed        int
	screenWidth  int
	screenHeight int
}

const (
	initZoom  = 10
	initSpeed = 10
)

func New(screenWidth int, screenHeight int, objects []*cp.Shape) *Universe {
	space := cp.NewSpace()
	//space.Iterations = 1

	for _, shape := range objects {
		space.AddBody(shape.Body())
		space.AddShape(shape)
	}

	return &Universe{
		space:        space,
		zoom:         initZoom,
		speed:        initSpeed,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}
}

func (u *Universe) Layout(_, _ int) (int, int) {
	return u.screenWidth, u.screenHeight
}

func (u *Universe) Update() error {
	for range u.speed {
		u.applyGravity()
		u.space.Step(1.0 / float64(ebiten.TPS()))
	}
	u.recordTrails()
	u.handleInput()
	return nil
}

func (u *Universe) applyGravity() {
	u.space.EachBody(func(body *cp.Body) {
		f := u.TotalGravitationalForce(u.space, body)
		body.SetForce(f)
	})
}

func (u *Universe) recordTrails() {
	u.space.EachBody(func(body *cp.Body) {
		body.UserData.(*object).addTrail(body.Position())
	})
}

func (u *Universe) handleInput() {
	for _, key := range inpututil.AppendJustPressedKeys(nil) {
		switch key {
		case ebiten.KeyArrowLeft:
			u.speed /= 2
			if u.speed < 1 {
				u.speed = 1
			}
		case ebiten.KeyArrowRight:
			u.speed *= 2
		case ebiten.KeyArrowUp:
			u.zoom /= 2
			if u.zoom < 1 {
				u.zoom = 1
			}
		case ebiten.KeyArrowDown:
			u.zoom *= 2
		}
	}
}

var (
	dot *ebiten.Image
)

func init() {
	dot = ebiten.NewImage(5, 5)
	vector.DrawFilledCircle(dot, 0, 0, 5, colornames.Blue, false)
}

func (u *Universe) Draw(screen *ebiten.Image) {
	var op ebiten.DrawImageOptions
	u.space.EachBody(func(body *cp.Body) {
		o := body.UserData.(*object)
		// find top left corner of the image
		r := o.image.Bounds()
		u.drawImage(screen, o.image, &op, body.Position().Sub(cp.Vector{X: float64(r.Dx() / 2), Y: float64(r.Dy() / 2)}))

		// draw the trails
		for _, trail := range o.trails {
			u.drawImage(screen, dot, &op, trail)
		}

		//if trailsImage, topLeft := o.drawTrails(colornames.Blue); trailsImage != nil {
		//	u.drawImage(screen, trailsImage, &op, topLeft, bodyOffset, 1/u.zoom, screenOffset)
		//}
	})

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f v: %s speed: %d, zoom: %.0f",
		ebiten.ActualTPS(),
		strings.Join(u.velocities(), ", "),
		u.speed,
		u.zoom,
	))
}

// drawImage draws the src image onto the dst image
//
// pos indicates the centre of the src image (in "universe" coordinates)
// offset is added to the translation from universe to screen coordinates. This allows the screen to stay centred on one object
func (u *Universe) drawImage(dst, src *ebiten.Image, op *ebiten.DrawImageOptions, pos cp.Vector) {
	op.GeoM = u.getGeometry(pos)
	dst.DrawImage(src, op)
}

func (u *Universe) getGeometry(pos cp.Vector) ebiten.GeoM {
	var centreObject cp.Vector
	if u.FocusObject != nil {
		centreObject = u.FocusObject.Body().Position().Mult(-1)
	}
	var geom ebiten.GeoM
	geom.Translate(pos.X, pos.Y)                                        // centre of the object
	geom.Translate(centreObject.X, centreObject.Y)                      // offset to keep focus on one object
	geom.Scale(1/u.zoom, 1/u.zoom)                                      // scale the image for the current zoom factor
	geom.Translate(float64(u.screenWidth/2), float64(u.screenHeight/2)) // screenOffset puts (0,0) at the centre of the screen
	return geom
}

func (u *Universe) velocities() []string {
	var velocities []string
	u.space.EachBody(func(body *cp.Body) {
		velocities = append(velocities, //strconv.FormatFloat(body.Velocity().Length(), 'f', 2, 64))
			fmt.Sprintf("%.2f(%d)", body.Velocity().Length(), len(body.UserData.(*object).trails)),
		)
	})
	return velocities
}

func NewBody(mass float64, radius float64, position cp.Vector, velocity cp.Vector, color color.Color) *cp.Shape {
	body := cp.NewBody(mass, cp.INFINITY)
	body.SetPosition(position)
	body.SetVelocity(velocity.X, velocity.Y)

	shape := cp.NewCircle(body, radius, cp.Vector{})
	shape.SetElasticity(1.0)
	shape.SetFriction(1)

	img := ebiten.NewImage(int(2*radius), int(2*radius))
	vector.DrawFilledCircle(img, float32(radius), float32(radius), float32(radius), color, false)
	body.UserData = newObject(img)

	return shape
}

const trails = 1000

type object struct {
	image  *ebiten.Image
	trails []cp.Vector
}

func newObject(image *ebiten.Image) *object {
	return &object{
		image:  image,
		trails: make([]cp.Vector, 0, trails),
	}
}

func (o *object) addTrail(position cp.Vector) {
	o.trails = append(o.trails, position)
	if len(o.trails) >= trails {
		o.trails = o.trails[1:]
	}
}
