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
	FocusObject  *cp.Shape
	space        *cp.Space
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
		f := gravity.TotalGravitationalForce(u.space, body)
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
	screenOffset := cp.Vector{X: float64(u.screenWidth / 2), Y: float64(u.screenHeight / 2)}
	u.space.EachBody(func(body *cp.Body) {
		var offset cp.Vector
		if u.FocusObject != nil {
			offset = u.FocusObject.Body().Position().Mult(-1)
		}
		o := body.UserData.(*object)
		drawImage(screen, o.image, &op, body.Position(), offset, 1/u.zoom, screenOffset)

		// draw the trails
		for _, trail := range o.trails {
			drawImage(screen, dot, &op, trail, offset, 1/u.zoom, screenOffset)
		}
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
// scale is the ratio between universe & screen coordinates. This allows to "zoom" in and out of the universe
// screenOffset should be half of screen width & height, putting the (0,0) universe coordinates in the middle of the screen
func drawImage(dst, src *ebiten.Image, op *ebiten.DrawImageOptions, pos, offset cp.Vector, scale float64, screenOffset cp.Vector) {
	// find top left corner of the image
	r := src.Bounds()
	offset = offset.Sub(cp.Vector{X: float64(r.Dx() / 2), Y: float64(r.Dy() / 2)})

	op.GeoM.Reset()
	op.GeoM.Translate(pos.X, pos.Y)                   // centre of the object
	op.GeoM.Translate(offset.X, offset.Y)             // offset to the top left corner of the image, possibly adapted to keep focus on one object
	op.GeoM.Scale(scale, scale)                       // scale the image for the current zoom factor
	op.GeoM.Translate(screenOffset.X, screenOffset.Y) // screenOffset puts (0,0) at the centre of the screen
	dst.DrawImage(src, op)
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
	shape.SetElasticity(0.1)
	shape.SetFriction(1)

	img := ebiten.NewImage(int(2*radius), int(2*radius))
	vector.DrawFilledCircle(img, float32(radius), float32(radius), float32(radius), color, false)
	body.UserData = newObject(img)

	return shape
}

const trails = 500

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
