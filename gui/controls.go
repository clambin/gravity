package gui

import (
	"fmt"
	"gravity/particle"
)

func (ui *UI) ClearObjects() {
	var particles []*particle.Particle
	for _, p := range ui.Space.Particles {
		if _, ok := ui.manualObjects[p]; ok == false {
			particles = append(particles, p)
		}
	}
	ui.Space.Particles = particles
	ui.manualObjects = make(map[*particle.Particle]struct{})
}

func (ui *UI) DecelerateObjects() {
	for _, p := range ui.Space.Particles {
		//if object.Manual == true {
		p.AX /= 2
		p.AY /= 2
		fmt.Printf("decelerate: (%f,%f)\n", p.AX, p.AY)
		//}
	}
}

func (ui *UI) AccelerateObjects() {
	for _, p := range ui.Space.Particles {
		// if object.Manual == true {
		p.AX *= 2
		p.AY *= 2
		fmt.Printf("accelerate: (%f,%f)\n", p.AX, p.AY)
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
	}, true)
}

func (ui *UI) toggleShowTrails() {
	ui.showTrails = !ui.showTrails
}
