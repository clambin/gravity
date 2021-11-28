package gui

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk"
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
		ui.ProcessEvents()

		ui.win.SetTitle(fmt.Sprintf("gravity (%.1f FPS)", 1.0/time.Now().Sub(timestamp).Seconds()))
		timestamp = time.Now()

		<-ticker.C
	}
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
