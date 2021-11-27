package gui

import (
	"fmt"
	"github.com/faiface/pixel"
	"gravity/particle"
)

func (ui *UI) ClearObjects() {
	var objects []*Object
	var particles []*particle.Particle
	for _, object := range ui.Objects {
		if object.Manual == false {
			objects = append(objects, object)
			particles = append(particles, object.particle)
		}
	}
	ui.Objects = objects
	ui.Space.Particles = particles
}

func (ui *UI) DecelerateObjects() {
	for _, object := range ui.Objects {
		//if object.Manual == true {
		object.particle.AX /= 2
		object.particle.AY /= 2
		fmt.Printf("decelerate: (%f,%f)\n", object.particle.AX, object.particle.AY)
		//}
	}
}

func (ui *UI) AccelerateObjects() {
	for _, object := range ui.Objects {
		// if object.Manual == true {
		object.particle.AX *= 2
		object.particle.AY *= 2
		fmt.Printf("accelerate: (%f,%f)\n", object.particle.AX, object.particle.AY)
		//}
	}
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
	VX := (position2.X - ui.position.X) / 50
	VY := (position2.Y - ui.position.Y) / 50
	ui.position = ui.Unscale(ui.position)
	ui.Add(&particle.Particle{
		M:  0.1,
		X:  ui.position.X,
		Y:  ui.position.Y,
		VX: VX,
		VY: VY,
	}, pixel.RGB(1, 1, 1), true)
}

func (ui *UI) toggleShowTrails() {
	ui.showTrails = !ui.showTrails
}
