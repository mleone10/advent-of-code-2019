package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type asteroid struct {
	x, y int     // Positive cartesian coordinates from an origin in the top-left of the field
	d, t float64 // Polar coordinates from the monitoring station
}

func main() {
	as := readInput()
	stationIndex := locateStation(as)

	log.Printf("Best location for station is at (%d,%d)", as[stationIndex].x, as[stationIndex].y)
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

func locateStation(as []asteroid) int {
	var optimalIndex, optimalScore int
	for i, a := range as {
		s := calculateScore(a, as)
		if s > optimalScore {
			optimalIndex = i
		}
	}
	return optimalIndex
}

func calculateScore(a asteroid, as []asteroid) int {
	return 0
}
