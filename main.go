package main

import (
	"github.com/clambin/gravity/internal/universe/universe"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp/v2"
	"golang.org/x/image/colornames"
	"log"
)

const (
	screenWidth  = 1600
	screenHeight = 800
)

func solarPlanets(_, _ float64) []*cp.Shape {
	const radius = 10
	initialVelocity := cp.Vector{X: 0, Y: -35}
	shapes := []*cp.Shape{
		universe.NewObject(1e7, 7*radius, cp.Vector{}, cp.Vector{}, colornames.Yellow),
		universe.NewObject(1, radius, cp.Vector{X: 400}, initialVelocity.Mult(2), colornames.Red),
		universe.NewObject(1e5, 4*radius, cp.Vector{X: -800}, initialVelocity.Mult(-1), colornames.Green),
	}
	for moonPos := 0; moonPos < 200; moonPos += 3 * radius {
		shapes = append(shapes, universe.NewObject(1, radius/2, cp.Vector{X: -600, Y: float64(-moonPos)}, initialVelocity.Mult(-1), colornames.Blue))
	}
	return shapes
}

func threeBodies(_, _ float64) []*cp.Shape {
	const radius = 10
	initialVelocity := cp.Vector{X: 0, Y: -35}
	return []*cp.Shape{
		universe.NewObject(1e7, 7*radius, cp.Vector{X: -400, Y: 1000}, cp.Vector{}, colornames.Yellow),
		universe.NewObject(1e7, 7*radius, cp.Vector{X: -1200, Y: 1000}, initialVelocity.Mult(-1), colornames.Red),
		universe.NewObject(1e7, 7*radius, cp.Vector{X: 800, Y: 1000}, initialVelocity.Mult(0.75), colornames.Blue),
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebiten")

	g := universe.New(screenWidth, screenHeight, threeBodies(screenHeight*5, screenHeight*5))
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
