package main

import (
	"bufio"
	"log"
	"os"
	"sort"
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

	log.Printf("Ore required for 1 fuel: %d", cs.simplify(fuel, 1, make(requirements)))
	log.Printf("Fuel produced from 1 trillion ore: %d", cs.maxFuel(1000000000000))
}

func (cs chemicals) simplify(c string, y int, r requirements) int {
	if c == ore {
		return y
	}

	if r[c] >= y {
		r[c] -= y
		return 0
	}

	if r[c] > 0 {
		y -= r[c]
		r[c] = 0
	}

	batches := aoc.Ceil(float64(y) / float64(cs[c].yield))
	oreSum := 0
	for k, amt := range cs[c].components {
		oreSum += cs.simplify(k, amt*batches, r)
	}

	r[c] += (batches * cs[c].yield) - y

	return oreSum
}

func (cs chemicals) maxFuel(oreMax int) int {
	return sort.Search(oreMax, func(i int) bool {
		return cs.simplify(fuel, i+1, make(requirements)) > oreMax
	})
}
