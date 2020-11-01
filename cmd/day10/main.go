package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
)

type location struct {
	x, y int
}

type view map[float64]interface{}

func main() {
	ls := readInput()
	l, v := locateStation(ls)

	log.Printf("Best location for station is at (%d,%d) with view to %d asteroids", l.x, l.y, len(v))
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
		v[math.Atan2(float64(y), float64(x))] = struct{}{}
	}
	return v
}
