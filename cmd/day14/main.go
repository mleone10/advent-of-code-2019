package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

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

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l := strings.Split(scanner.Text(), " => ")
		rs := make(requirements)
		for _, r := range strings.Split(l[0], ",") {
			ps := strings.Split(strings.TrimSpace(r), " ")
			y, _ := strconv.Atoi(ps[0])
			rs[ps[1]] = y
		}
		p := strings.Split(strings.TrimSpace(l[1]), " ")
		y, _ := strconv.Atoi(p[0])
		cs[p[1]] = chemical{yield: y, components: rs}
	}

	log.Printf("%+v", cs)
	log.Printf("Ore required for 1 fuel: %d", cs.calculateOre(cs.simplify(fuel, 1)))
}

func (cs chemicals) simplify(c string, y int) requirements {
	for k := range cs[c].components {
		if k == ore {
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
