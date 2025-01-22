package universe

import (
	"fmt"
	"github.com/clambin/gravity/gravity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/jakecoffman/cp/v2"
	"image/color"
	"strconv"
	"strings"
)

type Game struct {
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

func New(screenWidth int, screenHeight int, objects []*cp.Shape) *Game {
	space := cp.NewSpace()
	//space.Iterations = 1

	for _, shape := range objects {
		space.AddBody(shape.Body())
		space.AddShape(shape)
	}

	return &Game{
		space:        space,
		zoom:         initZoom,
		speed:        initSpeed,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return g.screenWidth, g.screenHeight
}

func (g *Game) Update() error {
	for range g.speed {
		g.applyGravity()
		g.space.Step(1.0 / float64(ebiten.TPS()))
	}
	g.handleInput()
	return nil
}

func (g *Game) applyGravity() {
	g.space.EachBody(func(body *cp.Body) {
		f := gravity.TotalGravitationalForce(g.space, body)
		body.SetForce(f)
	})
}

func (g *Game) handleInput() {
	for _, key := range inpututil.AppendJustPressedKeys(nil) {
		switch key {
		case ebiten.KeyArrowLeft:
			g.speed /= 2
			if g.speed < 1 {
				g.speed = 1
			}
		case ebiten.KeyArrowRight:
			g.speed *= 2
		case ebiten.KeyArrowUp:
			g.zoom /= 2
			if g.zoom < 1 {
				g.zoom = 1
			}
		case ebiten.KeyArrowDown:
			g.zoom *= 2
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	var op ebiten.DrawImageOptions
	g.space.EachBody(func(body *cp.Body) {
		img := body.UserData.(*ebiten.Image)
		r := img.Bounds()
		width, height := r.Dx(), r.Dy()

		op.GeoM.Reset()
		if g.FocusObject != nil {
			centre := g.FocusObject.Body().Position()
			op.GeoM.Translate(-centre.X, -centre.Y)
		}
		op.GeoM.Translate(-float64(width/2), -float64(height/2))
		op.GeoM.Translate(body.Position().X, body.Position().Y)
		op.GeoM.Scale(1/g.zoom, 1/g.zoom)
		op.GeoM.Translate(float64(g.screenWidth/2), float64(g.screenHeight/2))

		screen.DrawImage(img, &op)
	})

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f v: %s speed: %d, zoom: %.0f",
		ebiten.ActualTPS(),
		strings.Join(g.velocities(), ", "),
		g.speed,
		g.zoom,
	))
}

func (g *Game) velocities() []string {
	var velocities []string
	g.space.EachBody(func(body *cp.Body) {
		velocities = append(velocities, strconv.FormatFloat(body.Velocity().Length(), 'f', 2, 64))
	})
	return velocities
}

func NewObject(mass float64, radius float64, position cp.Vector, velocity cp.Vector, color color.Color) *cp.Shape {
	body := cp.NewBody(mass, cp.INFINITY)
	body.SetPosition(position)
	body.SetVelocity(velocity.X, velocity.Y)

	shape := cp.NewCircle(body, radius, cp.Vector{})
	shape.SetElasticity(0.1)
	shape.SetFriction(1)

	img := ebiten.NewImage(int(2*radius), int(2*radius))
	vector.DrawFilledCircle(img, float32(radius), float32(radius), float32(radius), color, false)
	body.UserData = img

	return shape
}
