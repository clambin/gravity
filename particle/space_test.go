package particle_test

import (
	"gravity/particle"
	"math/rand"
	"testing"
)

func BenchmarkSpace(b *testing.B) {
	s := particle.Space{Particles: make([]*particle.Particle, 0)}

	for i := 0; i < 10; i++ {
		s.Add(&particle.Particle{
			M: 1,
			X: float64(rand.Int31n(50) - 100),
			Y: float64(rand.Int31n(50) - 100),
		})
	}

	for i := 0; i < 10000; i++ {
		s.Step()
	}
}
