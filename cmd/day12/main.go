package main

import (
	"log"

	aoc "github.com/mleone10/advent-of-code-2019"
)

const (
	n    = 10
	dimX = iota
	dimY
	dimZ
)

type system struct {
	ms, init []moon
}

type moon struct {
	pos, vel vector
}

type vector struct {
	x, y, z int
}

func main() {
	s := newSystem([]vector{
		vector{0, 6, 1},
		vector{4, 4, 19},
		vector{-11, 1, 8},
		vector{2, 19, 15},
	})

	for i := 0; i < n; i++ {
		s.step()
	}

	log.Printf("Total energy after %d steps: %d", n, s.getEnergy())

	log.Printf("System-wide period: %d", s.getPeriod())
}

func newSystem(initPs []vector) system {
	ms := []moon{}
	initMs := []moon{}
	for _, p := range initPs {
		m := moon{p, vector{}}
		ms = append(ms, m)
		initMs = append(initMs, m)
	}

	return system{
		ms:   ms,
		init: initMs,
	}
}

func (s *system) reset() {
	for i, m := range s.init {
		s.ms[i] = m
	}
}

func (s *system) step() {
	s.applyGravity()
	s.applyVelocity()
}

func (s *system) getPeriod() int {
	s.reset()

	periods := [3]int{}
	done := [3]bool{}
	for !done[0] || !done[1] || !done[2] {
		s.step()
		for p := range periods {
			if !done[p] {
				equal := true
				for m := range s.ms {
					if !s.ms[m].equal(s.init[m], p) {
						equal = false
					}
				}
				if equal {
					done[p] = true
				} else {
					periods[p]++
				}
			}
		}
	}
	for i := range periods {
		periods[i]++
	}

	return aoc.Lcm(aoc.Lcm(periods[0], periods[1]), periods[2])
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

func (m *moon) equal(n moon, p int) bool {
	if p == dimX {
		return m.pos.x == n.pos.x && m.vel.x == n.vel.x
	} else if p == dimY {
		return m.pos.y == n.pos.y && m.vel.y == n.vel.y
	} else {
		return m.pos.z == n.pos.z && m.vel.z == n.vel.z
	}
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
