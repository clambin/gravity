package gui

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
	"golang.org/x/image/colornames"
	"math"
	"time"
)

type UI struct {
	X             float64
	Y             float64
	Space         *chipmunk.Space
	manualObjects map[*chipmunk.Body]struct{}
	trails        map[*chipmunk.Body][]pixel.Vec
	scale         float64
	time          int
	showTrails    bool
	position      pixel.Vec
	win           *pixelgl.Window
}

func NewUI(X, Y float64) (ui *UI) {
	return &UI{
		X:             X,
		Y:             Y,
		Space:         chipmunk.NewSpace(),
		manualObjects: make(map[*chipmunk.Body]struct{}),
		trails:        make(map[*chipmunk.Body][]pixel.Vec),
		scale:         1,
		time:          1,
	}
}

func (ui UI) Scale(input pixel.Vec) pixel.Vec {
	return pixel.Vec{
		X: input.X / ui.scale,
		Y: input.Y / ui.scale,
	}
}

func (ui UI) Unscale(input pixel.Vec) pixel.Vec {
	return pixel.Vec{
		X: input.X * ui.scale,
		Y: input.Y * ui.scale,
	}
}

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
	imd.Push(ui.Scale(pixel.V(float64(p.X), float64(p.Y))))
	imd.Circle(float64(body.Shapes[0].GetAsCircle().Radius)/ui.scale, 0)
}

func (ui *UI) drawTrail(imd *imdraw.IMDraw, body *chipmunk.Body) {
	for index := range ui.trails[body] {
		if index == 0 {
			continue
		}
		imd.Color = pixel.RGB(0, 0, 0.5)
		imd.Push(ui.Scale(ui.trails[body][index-1]))
		imd.Push(ui.Scale(ui.trails[body][index]))
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
	imd.Push(ui.Scale(pixel.V(x, y)), ui.Scale(pixel.V(x+vx, y+vy)))
	imd.Line(1)
}

func (ui *UI) drawAcceleration(imd *imdraw.IMDraw, body *chipmunk.Body) {
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

func (ui *UI) Add(x, y, r, m, vx, vy float64, manual bool) {
	circle := chipmunk.NewCircle(vect.Vector_Zero, float32(r))
	circle.SetElasticity(0.95)
	body := chipmunk.NewBody(vect.Float(m), circle.Moment(float32(m)))
	body.SetPosition(vect.Vect{X: vect.Float(x), Y: vect.Float(y)})
	body.AddShape(circle)
	body.SetVelocity(float32(vx), float32(vy))
	ui.Space.AddBody(body)

	if manual {
		ui.manualObjects[body] = struct{}{}
	}
}

func (ui *UI) RunGUI() {
	cfg := pixelgl.WindowConfig{
		Title:  "gravity",
		Bounds: pixel.R(-ui.X/2, -ui.Y/2, ui.X/2, ui.Y/2),
	}
	var err error
	ui.win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(40 * time.Millisecond)
	timestamp := time.Now()
	for !ui.win.Closed() {
		for i := 0; i < ui.time; i++ {
			ui.Space.Step(1.0)
			ui.gravity()
		}
		ui.recordTrails()
		ui.Draw(ui.win)
		ui.win.Update()
		ui.processEvents()

		ui.win.SetTitle(fmt.Sprintf("gravity (%.1f FPS)", 1.0/time.Now().Sub(timestamp).Seconds()))
		timestamp = time.Now()

		<-ticker.C
	}
}

func (ui *UI) gravity() {
	for _, body := range ui.Space.Bodies {
		ui.applyGravity(body)
	}
}

func distance(a, b *chipmunk.Body) float32 {
	dx := float64(a.Position().X - b.Position().X)
	dy := float64(a.Position().Y - b.Position().Y)
	return float32(math.Sqrt(dx*dx + dy*dy))
}

func (ui *UI) applyGravity(body *chipmunk.Body) {
	var fx, fy float32
	for _, other := range ui.Space.Bodies {
		if other == body {
			continue
		}

		dx := float32(other.Position().X - body.Position().X)
		dy := float32(other.Position().Y - body.Position().Y)
		r := distance(body, other)
		// F = m * a
		// => a = F / m
		// F1 = F2 = G * m1 * m2 / r^2
		// => a1 = G * m2 / r2
		//
		// ax = dx * a / r
		// => ax = dx * G * m2 / r^3
		r3 := r * r * r
		const G = 0.0000001
		fx += dx * G * (float32(other.Mass() * body.Mass())) / r3
		fy += dy * G * (float32(other.Mass() * body.Mass())) / r3
	}
	body.AddForce(fx, fy)
}

func (ui *UI) recordTrails() {
	const maxTrails = 2000
	for _, body := range ui.Space.Bodies {
		trails, _ := ui.trails[body]
		p := body.Position()
		trails = append(trails, pixel.V(float64(p.X), float64(p.Y)))
		if len(trails) > maxTrails {
			trails = trails[1:]
		}
		ui.trails[body] = trails
	}
}

func (ui *UI) processEvents() {
	if ui.win.JustPressed(pixelgl.KeyC) {
		ui.ClearObjects()
	}
	if ui.win.JustPressed(pixelgl.KeyD) {
		ui.DecelerateObjects()
	}
	if ui.win.JustPressed(pixelgl.KeyA) {
		ui.AccelerateObjects()
	}
	if ui.win.JustPressed(pixelgl.KeyT) {
		ui.toggleShowTrails()
	}
	if ui.win.JustReleased(pixelgl.KeyRight) {
		ui.SpeedUp()
	}
	if ui.win.JustReleased(pixelgl.KeyLeft) {
		ui.SlowDown()
	}
	if ui.win.JustPressed(pixelgl.MouseButtonLeft) {
		ui.AddObjectStart()
	}
	if ui.win.JustReleased(pixelgl.MouseButtonLeft) {
		ui.AddObject()
	}
	if scroll := ui.win.MouseScroll(); scroll.Y != 0 {
		ui.scale = math.Max(ui.scale-scroll.Y, 0.1)
		fmt.Printf("scale: %.1f\n", ui.scale)
	}
}
