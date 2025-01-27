package gravity

import (
	"github.com/jakecoffman/cp/v2"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
	"reflect"
	"testing"
)

func TestTotalGravitationalForce(t *testing.T) {
	space := cp.NewSpace()
	space.Iterations = 1

	const radius = 25
	var v cp.Vector
	shapes := []*cp.Shape{
		NewObject(10e1, radius, cp.Vector{X: 0, Y: 0}, cp.Vector{}, colornames.White),
		NewObject(10e3, radius, cp.Vector{X: 1000, Y: 0}, v.Mult(-1), colornames.Red),
		NewObject(10e1, radius, cp.Vector{X: 2000, Y: 0}, v.Mult(-1), colornames.Green),
	}
	for _, shape := range shapes {
		space.AddBody(shape.Body())
		space.AddShape(shape)
	}

	if got := TotalGravitationalForce(space, shapes[0].Body()); !reflect.DeepEqual(got, cp.Vector{X: 0.10025}) {
		t.Errorf("TotalGravitationalForce shape 0: got %v want %v", got, cp.Vector{X: 0.10025})
	}
	if got := TotalGravitationalForce(space, shapes[1].Body()); !reflect.DeepEqual(got, cp.Vector{}) {
		t.Errorf("TotalGravitationalForce shape 1: got %v want %v", got, cp.Vector{})
	}
	if got := TotalGravitationalForce(space, shapes[2].Body()); !reflect.DeepEqual(got, cp.Vector{X: -0.10025}) {
		t.Errorf("TotalGravitationalForce shape 2: got %v want %v", got, cp.Vector{X: -0.10025})
	}
}

func NewObject(mass float64, radius float64, position cp.Vector, velocity cp.Vector, _ color.Color) *cp.Shape {
	body := cp.NewBody(mass, cp.INFINITY)
	body.SetPosition(position)
	body.SetVelocity(velocity.X, velocity.Y)

	shape := cp.NewCircle(body, radius, cp.Vector{})
	shape.SetElasticity(0.1)
	shape.SetFriction(0.9)

	return shape
}

func Test_gravitationalForce(t *testing.T) {
	// horizontal
	sun := cp.NewBody(1.989e30, cp.INFINITY)             // sun mass is 1.989×10³⁰ kg
	sun.SetPosition(cp.Vector{X: 149_600_000_000, Y: 0}) // distance is 149,600,000 km
	earth := cp.NewBody(5.972e24, cp.INFINITY)           // earth mass is 5.972*10^24 kg
	f := gravitationalForce(earth, sun, 6.674e-11)
	// force should be 3.54×10²² N
	f.X = math.Round(100*f.X/1e22) / 100
	want := cp.Vector{X: 3.54, Y: 0}
	if !reflect.DeepEqual(want, f) {
		t.Errorf("gravitationalForce got %v want %v", f, want)
	}

	// vertical
	sun.SetPosition(cp.Vector{X: 0, Y: 149_600_000_000})
	f = gravitationalForce(earth, sun, 6.674e-11)
	f.Y = math.Round(100*f.Y/1e22) / 100
	want = cp.Vector{X: 0, Y: 3.54}
	if !reflect.DeepEqual(want, f) {
		t.Errorf("gravitationalForce got %v want %v", f, want)
	}
}
