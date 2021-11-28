package gui

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
	"math"
)

func (ui *UI) ProcessEvents() {
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
	if ui.win.JustReleased(pixelgl.Key0) {
		ui.scale = 1.0
		fmt.Printf("scale: %.1f\n", ui.scale)
	}
	if scroll := ui.win.MouseScroll(); scroll.Y != 0 {
		ui.scale = math.Max(ui.scale-scroll.Y, 0.1)
		fmt.Printf("scale: %.1f\n", ui.scale)
	}
}

func (ui *UI) ClearObjects() {
	var bodies []*chipmunk.Body
	for _, body := range ui.Space.Bodies {
		if _, ok := ui.manualObjects[body]; ok == false {
			bodies = append(bodies, body)
		}
	}
	ui.Space.Bodies = bodies
	ui.manualObjects = make(map[*chipmunk.Body]struct{})
}

func (ui *UI) DecelerateObjects() {
	/*
		for _, p := range ui.Space.Particles {
			//if object.Manual == true {
			p.AX /= 2
			p.AY /= 2
			fmt.Printf("decelerate: (%f,%f)\n", p.AX, p.AY)
			//}
		}
	*/
}

func (ui *UI) AccelerateObjects() {
	/*
		for _, p := range ui.Space.Particles {
			// if object.Manual == true {
			p.AX *= 2
			p.AY *= 2
			fmt.Printf("accelerate: (%f,%f)\n", p.AX, p.AY)
			//}
		}
	*/
}

func (ui *UI) SpeedUp() {
	ui.time *= 2
	fmt.Printf("time: %d\n", ui.time)
}

func (ui *UI) SlowDown() {
	if ui.time > 1 {
		ui.time /= 2
	}
	fmt.Printf("time: %d\n", ui.time)
}

func (ui *UI) AddObjectStart() {
	ui.position = ui.win.MousePosition()
}

func (ui *UI) AddObject() {
	fmt.Printf("adding particle at (%f,%f)\n", ui.position.X, ui.position.Y)
	position2 := ui.win.MousePosition()
	VX := float32(position2.X-ui.position.X) / 50
	VY := float32(position2.Y-ui.position.Y) / 50
	position := ui.Unscale(ui.position)
	ui.Add(vect.Vect{X: vect.Float(position.X), Y: vect.Float(position.Y)}, 1, 1, VX, VY, true)
}

func (ui *UI) Add(position vect.Vect, r, m, vx, vy float32, manual bool) {
	circle := chipmunk.NewCircle(vect.Vector_Zero, r)
	//circle.SetElasticity(0.1)
	body := chipmunk.NewBody(vect.Float(m), circle.Moment(m))
	body.SetPosition(position)
	body.AddShape(circle)
	body.SetVelocity(vx, vy)
	ui.Space.AddBody(body)

	if manual {
		ui.manualObjects[body] = struct{}{}
	}
}

func (ui *UI) toggleShowTrails() {
	ui.showTrails = !ui.showTrails
}
