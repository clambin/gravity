package gui

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"gravity/particle"
	"math"
	"time"
)

type UI struct {
	X             float64
	Y             float64
	Space         particle.Space
	manualObjects map[*particle.Particle]struct{}
	trails        map[*particle.Particle][]pixel.Vec
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
		Space:         particle.Space{Particles: make([]*particle.Particle, 0)},
		manualObjects: make(map[*particle.Particle]struct{}),
		trails:        make(map[*particle.Particle][]pixel.Vec),
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
	for _, p := range ui.Space.Particles {
		ui.drawObject(imd, p)
		if ui.showTrails {
			ui.drawTrail(imd, p)
		}
		ui.drawVelocity(imd, p)
		ui.drawAcceleration(imd, p)
	}
	imd.Draw(win)
}

func (ui *UI) drawObject(imd *imdraw.IMDraw, p *particle.Particle) {
	if _, ok := ui.manualObjects[p]; ok == true {
		imd.Color = colornames.White
	} else {
		imd.Color = colornames.Purple
	}
	imd.Push(ui.Scale(pixel.V(p.X, p.Y)))
	imd.Circle(2, 0)
}

func (ui *UI) drawTrail(imd *imdraw.IMDraw, p *particle.Particle) {
	for index := range ui.trails[p] {
		if index == 0 {
			continue
		}
		imd.Color = pixel.RGB(0, 0, 0.5)
		imd.Push(ui.Scale(ui.trails[p][index-1]))
		imd.Push(ui.Scale(ui.trails[p][index]))
		imd.Line(1)
	}
}

func (ui *UI) drawVelocity(imd *imdraw.IMDraw, p *particle.Particle) {
	x1 := p.X
	y1 := p.Y
	x2 := x1 + p.VX*10
	y2 := y1 + p.VY*10
	imd.Color = pixel.RGB(0, 0.8, 0)
	imd.Push(ui.Scale(pixel.V(x1, y1)), ui.Scale(pixel.V(x2, y2)))
	imd.Line(1)
}

func (ui *UI) drawAcceleration(imd *imdraw.IMDraw, p *particle.Particle) {
	x1 := p.X
	y1 := p.Y
	x2 := x1 + p.AX*1000
	y2 := y1 + p.AY*1000
	imd.Color = pixel.RGB(1, 0, 0)
	imd.Push(ui.Scale(pixel.V(x1, y1)), ui.Scale(pixel.V(x2, y2)))
	imd.Line(1)
}

func (ui *UI) Add(p *particle.Particle, manual bool) {
	ui.Space.Add(p)
	if manual {
		ui.manualObjects[p] = struct{}{}
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
		ui.processEvents()

		for i := 0; i < ui.time; i++ {
			ui.Space.Step()
		}
		ui.Draw(ui.win)
		ui.recordTrails()

		ui.win.Update()

		ui.win.SetTitle(fmt.Sprintf("gravity (%.1f FPS)", 1.0/time.Now().Sub(timestamp).Seconds()))
		timestamp = time.Now()

		<-ticker.C
	}
}

func (ui *UI) recordTrails() {
	const maxTrails = 2000
	for _, p := range ui.Space.Particles {
		trails, _ := ui.trails[p]
		trails = append(trails, pixel.V(p.X, p.Y))
		if len(trails) > maxTrails {
			trails = trails[1:]
		}
		ui.trails[p] = trails
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
