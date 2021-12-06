package gui

import (
	"fmt"
	"github.com/clambin/gravity/gui/field"
	"github.com/clambin/gravity/gui/stats"
	"github.com/clambin/gravity/gui/status"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk/vect"
	"golang.org/x/image/colornames"
	"image/color"
	"time"
)

type UI struct {
	X        float64
	Y        float64
	Field    *field.Field
	Stats    stats.Reporter
	Status   status.Reporter
	Callback CallbackFunc
	name     string
	time     int
	position pixel.Vec
}

type CallbackFunc func(ui *UI)

func NewUI(name string, X, Y float64) (ui *UI) {
	return &UI{
		X:     X,
		Y:     Y,
		Field: field.New(name, X, Y),
		name:  name,
		time:  1,
	}
}

func (ui *UI) RunGUI() {
	cfg := pixelgl.WindowConfig{
		Title:  ui.name,
		Bounds: pixel.R(-ui.X/2, -ui.Y/2, ui.X/2, ui.Y/2),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	ui.Field.MakeCanvas()
	ui.Stats.MakeCanvas(ui.X, ui.Y)
	ui.Status.MakeCanvas(ui.X, ui.Y)

	ticker := time.NewTicker(40 * time.Millisecond)
	timestamp := time.Now()
	for !win.Closed() {
		ui.Field.Steps(ui.time)
		if ui.Callback != nil {
			ui.Callback(ui)
		}

		win.Clear(colornames.Black)
		ui.Field.Draw(win)
		ui.Status.Draw(win, ui.Field.ViewFinder.Offset, ui.Field.ViewFinder.Scale, ui.time)
		ui.Stats.Draw(win, ui.Field.Stats())

		win.Update()
		ui.ProcessEvents(win)

		win.SetTitle(fmt.Sprintf("%s (%.1f FPS)", ui.name, 1.0/time.Now().Sub(timestamp).Seconds()))
		timestamp = time.Now()

		<-ticker.C
	}
}

// SetTime sets the speedup factor for the simulation
func (ui *UI) SetTime(factor int) {
	ui.time = factor
}

type Body struct {
	Position vect.Vect
	Radius   float32
	Mass     vect.Float
	Velocity vect.Vect
	Color    color.Color
}

func (ui *UI) Load(bodies []Body) {
	for _, body := range bodies {
		ui.Field.Add(body.Position, body.Radius, body.Mass, body.Velocity, false)
	}
}

func (ui *UI) AddManual(position pixel.Vec, radius float32, mass float32, velocity pixel.Vec) {
	ui.Field.Add(
		ui.Field.ViewFinder.ViewFinderToReal(position),
		radius,
		vect.Float(mass),
		ui.Field.ViewFinder.ViewFinderToReal(velocity),
		true)
}
