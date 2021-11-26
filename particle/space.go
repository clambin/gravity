package particle

type Space struct {
	Particles []*Particle
}

func (s *Space) Step() {
	for _, p := range s.Particles {
		p.Update(s.Particles)
	}
}

func (s *Space) Add(p *Particle) {
	s.Particles = append(s.Particles, p)
}
