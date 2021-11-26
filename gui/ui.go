package gui

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"gravity/particle"
	"math"
	"time"
)

type UI struct {
	X        float64
	Y        float64
	Objects  []*Object
	Space    particle.Space
	scale    float64
	time     int
	position pixel.Vec
	actions  []action
	win      *pixelgl.Window
}

type action struct {
	button   pixelgl.Button
	command  func()
	released bool
}

func NewUI(X, Y float64) (ui *UI) {
	ui = &UI{
		X:       X,
		Y:       Y,
		Objects: make([]*Object, 0),
		Space: particle.Space{
			Particles: make([]*particle.Particle, 0),
		},
		scale: 1,
		time:  1,
	}
	ui.actions = []action{
		{button: pixelgl.KeyC, command: ui.ClearObjects},
		{button: pixelgl.KeyD, command: ui.DecelerateObjects},
		{button: pixelgl.KeyA, command: ui.AccelerateObjects},
		{button: pixelgl.KeyRight, command: ui.SpeedUp},
		{button: pixelgl.KeyLeft, command: ui.SlowDown},
		{button: pixelgl.MouseButtonLeft, command: ui.AddObjectStart},
		{button: pixelgl.MouseButtonLeft, command: ui.AddObject, released: true},
	}

	return
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
	for _, o := range ui.Objects {
		o.Draw(win)
	}
	win.Update()
}

func (ui *UI) Add(p *particle.Particle, color pixel.RGBA, manual bool) {
	ui.Space.Add(p)
	ui.Objects = append(ui.Objects, NewObject(p, color, ui, manual))
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

	for !ui.win.Closed() {
		ui.processEvents()

		ui.Space.Step()
		ui.Draw(ui.win)
		ui.recordTrails()

		time.Sleep(time.Duration(100000/ui.time) * time.Microsecond)
	}
}

func (ui *UI) recordTrails() {
	for _, o := range ui.Objects {
		o.recordTrail()
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
