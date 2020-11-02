package main

import (
	"log"

	aoc "github.com/mleone10/advent-of-code-2019"
)

const n = 1000

type system struct {
	ms []moon
}

type moon struct {
	pos, vel vector
}

type vector struct {
	x, y, z int
}

func main() {
	s := system{
		ms: []moon{
			moon{vector{0, 6, 1}, vector{}},
			moon{vector{4, 4, 19}, vector{}},
			moon{vector{-11, 1, 8}, vector{}},
			moon{vector{2, 19, 15}, vector{}},
		},
	}

	for i := 0; i < n; i++ {
		s.step()
	}

	log.Printf("Total energy after %d steps: %d", n, s.getEnergy())
}

func (s *system) step() {
	s.applyGravity()
	s.applyVelocity()
}

func (s *system) applyGravity() {
	for i := range s.ms {
		for j, n := range s.ms {
			if i == j {
				continue
			}
			s.ms[i].applyGravity(n)
		}
	}
}

func (s *system) applyVelocity() {
	for i := range s.ms {
		s.ms[i].applyVelocity()
	}
}

func (s *system) getEnergy() int {
	var sum int
	for _, m := range s.ms {
		sum += m.getEnergy()
	}
	return sum
}

func (m *moon) applyGravity(n moon) {
	m.vel.applyGravity(m.pos, n.pos)
}

func (m *moon) applyVelocity() {
	m.pos.add(m.vel)
}

func (m *moon) getEnergy() int {
	return m.getPotentialEnergy() * m.getKineticEnergy()
}

func (m *moon) getPotentialEnergy() int {
	return m.pos.sum()
}

func (m *moon) getKineticEnergy() int {
	return m.vel.sum()
}

func (v *vector) applyGravity(m, n vector) {
	gravityDelta := func(a, b int) int {
		if a > b {
			return -1
		}
		if a == b {
			return 0
		}
		return 1
	}
	v.x += gravityDelta(m.x, n.x)
	v.y += gravityDelta(m.y, n.y)
	v.z += gravityDelta(m.z, n.z)
}

func (v *vector) add(r vector) {
	v.x += r.x
	v.y += r.y
	v.z += r.z
}

func (v *vector) sum() int {
	return aoc.Abs(v.x) + aoc.Abs(v.y) + aoc.Abs(v.z)
}
