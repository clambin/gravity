package gui

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"gravity/gui/field"
	"gravity/gui/stats"
	"image/color"
	"time"
)

type UI struct {
	X        float64
	Y        float64
	Field    *field.Field
	Stats    stats.Stats
	Callback CallbackFunc
	time     int
	position pixel.Vec
}

type CallbackFunc func(ui *UI)

func NewUI(X, Y float64) (ui *UI) {
	return &UI{
		X:     X,
		Y:     Y,
		Field: field.New(),
		time:  1,
	}
}

func (ui *UI) RunGUI() {
	cfg := pixelgl.WindowConfig{
		Title:  "gravity",
		Bounds: pixel.R(-ui.X/2, -ui.Y/2, ui.X/2, ui.Y/2),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	ui.Field.MakeCanvas(ui.X, ui.Y)
	ui.Stats.MakeCanvas(ui.X, ui.Y)

	ticker := time.NewTicker(40 * time.Millisecond)
	timestamp := time.Now()
	for !win.Closed() {
		ui.Field.Steps(ui.time)

		win.Clear(colornames.Black)

		ui.Field.Draw(win)
		ui.Stats.Draw(win, ui.Field.Stats())

		win.Update()
		ui.ProcessEvents(win)

		if ui.Callback != nil {
			ui.Callback(ui)
		}

		win.SetTitle(fmt.Sprintf("gravity (%.1f FPS)", 1.0/time.Now().Sub(timestamp).Seconds()))
		timestamp = time.Now()

		<-ticker.C
	}
}

type Body struct {
	X     float64
	Y     float64
	R     float32
	M     float32
	VX    float64
	VY    float64
	Color color.Color
}

func (ui *UI) Load(bodies []Body) {
	for _, body := range bodies {
		ui.Field.Add(pixel.V(body.X, body.Y), body.R, body.M, pixel.V(body.VX, body.VY), false)
	}
}
