package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strings"

	aoc "github.com/mleone10/advent-of-code-2019"
)

type location struct {
	x, y int
	r    float64
}

type view map[float64][]location

func main() {
	l, v := locateStation(readInput())
	log.Printf("Best location for station is at (%d,%d) with view to %d asteroids", l.x, l.y, len(v))

	a := v.findNthDestroyed(200)
	log.Printf("200th asteroid to be destroyed is at (%d,%d), whose hash is %d", a.x, a.y, a.x*100+a.y)
}

func readInput() []location {
	ls := []location{}

	scanner := bufio.NewScanner(os.Stdin)
	y := 0
	for scanner.Scan() {
		x := 0
		for _, l := range strings.Split(scanner.Text(), "") {
			if l == "#" {
				ls = append(ls, location{x: x, y: y})
			}
			x++
		}
		y++
	}

	return ls
}

func locateStation(ls []location) (location, view) {
	var optLoc location
	var optView view
	for _, l := range ls {
		v := computeStationView(l, ls)
		if len(v) > len(optView) {
			optView = v
			optLoc = l
		}
	}
	return optLoc, optView
}

func computeStationView(s location, ls []location) view {
	v := view{}
	for _, l := range ls {
		x, y := l.x-s.x, l.y-s.y
		a := aoc.Bearing(x, y)
		l.r = math.Hypot(float64(x), float64(y))
		if l.r > 0 {
			v[a] = append(v[a], l)
		}
	}

	for _, a := range v {
		sort.Slice(a, func(i, j int) bool {
			return a[i].r < a[j].r
		})
	}

	return v
}

func (v view) findNthDestroyed(n int) location {
	var l location
	for _, a := range angles(v) {
		l, v[a] = v[a][0], v[a][1:]
		if len(v[a]) == 0 {
			delete(v, a)
		}
		n--
		if n <= 0 {
			break
		}
	}
	return l
}

func angles(v view) []float64 {
	as := []float64{}
	for a := range v {
		as = append(as, a)
	}
	sort.Float64s(as)
	return as
}
