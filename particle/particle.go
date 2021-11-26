package particle

import "math"

type Particle struct {
	M  float64
	X  float64
	Y  float64
	VX float64
	VY float64
	AX float64
	AY float64
}

const (
	G = 0.001
)

func (p *Particle) Update(particles []*Particle) {
	p.updatePosition()
	p.updateVelocity()
	p.updateAcceleration(particles)
}

func (p *Particle) updatePosition() {
	p.X += p.VX
	p.Y += p.VY
}

func (p *Particle) updateVelocity() {
	p.VX += p.AX
	p.VY += p.AY
}

func (p *Particle) updateAcceleration(particles []*Particle) {
	p.AX = 0
	p.AY = 0
	for _, particle := range particles {
		if particle == p {
			continue
		}
		p.gravitation(particle)
	}
}

func (p Particle) distance(other *Particle) (distance float64) {
	dx := other.X - p.X
	dy := other.Y - p.Y
	distance = math.Sqrt(dx*dx + dy*dy)
	if distance < 1 {
		distance = 1
	}
	return
}

func (p *Particle) gravitation(other *Particle) {
	dx := other.X - p.X
	dy := other.Y - p.Y
	r := p.distance(other)
	// F = m * a
	// => a = F / m
	// F1 = F2 = G * m1 * m2 / r^2
	// => a1 = G * m2 / r2
	//
	// ax = dx * a / r
	// => ax = dx * G * m2 / r^3
	r3 := r * r * r
	p.AX += dx * G * other.M / r3
	p.AY += dy * G * other.M / r3
}
