package main

import (
	"log"

	aoc "github.com/mleone10/advent-of-code-2019"
)

const (
	fuel = "FUEL"
	ore  = "ORE"
)

type chemical struct {
	yield      int
	components map[string]int
}

type chemicals map[string]chemical
type requirements map[string]int

func main() {
	cs := make(chemicals)

	// TODO: Populate chemicals map programatically
	cs["A"] = chemical{yield: 10, components: requirements{"ORE": 10}}
	cs["B"] = chemical{yield: 1, components: requirements{"ORE": 1}}
	cs["C"] = chemical{yield: 1, components: requirements{"A": 7, "B": 1}}
	cs["D"] = chemical{yield: 1, components: requirements{"A": 7, "C": 1}}
	cs["E"] = chemical{yield: 1, components: requirements{"A": 7, "D": 1}}
	cs["FUEL"] = chemical{yield: 1, components: requirements{"A": 7, "E": 1}}

	log.Printf("Ore required for 1 fuel: %d", cs.calculateOre(cs.simplify(fuel, 1)))
}

func (cs chemicals) simplify(c string, y int) requirements {
	for k := range cs[c].components {
		if k == ore {
			log.Println("Basic", c, y)
			return requirements{c: y}
		}
	}

	reqs := make(requirements)
	for k := range cs[c].components {
		s := cs.simplify(k, y*cs[c].components[k])
		for k, a := range s {
			reqs[k] += a
		}
	}

	return reqs
}

func (cs chemicals) calculateOre(rs requirements) int {
	var sum int
	for k := range rs {
		sum += aoc.Ceil(float64(rs[k])/float64(cs[k].yield)) * cs[k].yield
	}
	return sum
}
