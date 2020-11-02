package main

import "log"

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

func (s system) step() {
	s.applyGravity()
	s.applyVelocity()
}

func (s system) applyGravity() {

}

func (s system) applyVelocity() {
	for _, m := range s.ms {
		m.applyVelocity()
	}
}

func (s system) getEnergy() int {
	var sum int
	for _, m := range s.ms {
		sum += m.getEnergy()
	}
	return sum
}

func (m moon) applyVelocity() {
	m.pos.add(m.vel)
}

func (m moon) getEnergy() int {
	return m.getPotentialEnergy() * m.getKineticEnergy()
}

func (m moon) getPotentialEnergy() int {
	return m.pos.sum()
}

func (m moon) getKineticEnergy() int {
	return m.vel.sum()
}

func (v vector) add(r vector) {
	v.x += r.x
	v.y += r.y
	v.z += r.z
}

func (v vector) sum() int {
	return v.x + v.y + v.z
}
