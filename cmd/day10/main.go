package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type asteroid struct {
	x, y int     // Positive cartesian coordinates from an origin in the top-left of the field
	d, a float64 // Polar coordinates from the monitoring station
	view map[float64][]asteroid
}

func main() {
	as := readInput()
	a := locateStation(as)

	log.Printf("Best location for station is at (%d,%d)", a.x, a.y)
}

func readInput() []asteroid {
	as := []asteroid{}

	scanner := bufio.NewScanner(os.Stdin)
	y := 0
	for scanner.Scan() {
		x := 0
		for _, l := range strings.Split(scanner.Text(), "") {
			if l == "#" {
				as = append(as, asteroid{x: x, y: y})
			}
			x++
		}
		y++
	}

	return as
}

func locateStation(as []asteroid) asteroid {
	for i, a := range as {
		a.view = constructView(i, as)
	}

	var optimalAsteroid asteroid
	var optimalScore int
	for _, a := range as {
		s := len(a.view)
		if s > optimalScore {
			optimalAsteroid = a
		}
	}
	return optimalAsteroid
}

func constructView(originIndex int, as []asteroid) map[float64][]asteroid {
	view := map[float64][]asteroid{}
	originAsteroid := as[originIndex]

	for i, a := range as {
		if i != originIndex {

		}
	}

	return view
}
