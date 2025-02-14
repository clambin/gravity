package main

import (
	"github.com/clambin/gravity/internal/universe/universe"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp/v2"
	"golang.org/x/image/colornames"
	"log"
	"net/http"
	_ "net/http/pprof"
)

const (
	screenWidth  = 1600
	screenHeight = 800
)

var (
	solarPlanets []*cp.Shape
	threeBodies  []*cp.Shape
)

func main() {
	go func() {
		_ = http.ListenAndServe("localhost:6000", nil)
	}()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebiten")

	objects := threeBodies
	g := universe.New(screenWidth, screenHeight, objects)
	// use a fake gravitational constant so masses, distances & velocities are easier
	g.Gravity.G = 0.0674
	g.FocusObject = objects[0]
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func init() {
	solarPlanets = makeSolarPlanets()
	threeBodies = makeThreeBodies()
}

func makeSolarPlanets() []*cp.Shape {
	const radius = 10
	initialVelocity := cp.Vector{X: 0, Y: -35}
	shapes := []*cp.Shape{
		universe.NewBody(1e7, 7*radius, cp.Vector{}, cp.Vector{}, colornames.Yellow),
		universe.NewBody(1, radius, cp.Vector{X: 400}, initialVelocity.Mult(1.5), colornames.Red),
		universe.NewBody(1e5, 4*radius, cp.Vector{X: -800}, initialVelocity.Mult(-1), colornames.Green),
	}
	for moonPos := 0; moonPos < 400; moonPos += 3 * radius {
		shapes = append(shapes, universe.NewBody(1, radius/2, cp.Vector{X: -600, Y: float64(-moonPos)}, initialVelocity.Mult(-1), colornames.Grey))
	}
	return shapes
}

func makeThreeBodies() []*cp.Shape {
	const radius = 100
	initialVelocity := cp.Vector{X: 0, Y: -35}
	return []*cp.Shape{
		universe.NewBody(1e7, radius, cp.Vector{X: -400, Y: 1000}, cp.Vector{}, colornames.Yellow),
		universe.NewBody(1e7, radius, cp.Vector{X: -1200, Y: 1000}, initialVelocity.Mult(-1.5), colornames.Red),
		universe.NewBody(1e7, radius, cp.Vector{X: 800, Y: 1000}, initialVelocity.Mult(0.4), colornames.Blue),
	}
}
