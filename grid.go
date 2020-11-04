package aoc

import (
	"fmt"
	"strconv"
)

// Grid represents a two dimensional array of integers.
type Grid struct {
	field                  map[Coordinate]int
	minX, minY, maxX, maxY int
}

// Coordinate represents a two-dimensional (x,y) position on the grid.
type Coordinate struct {
	X, Y int
}

// Set stores integer i at location (x, y)
func (g *Grid) Set(x, y, i int) {
	if g.field == nil {
		g.field = map[Coordinate]int{}
	}

	g.field[Coordinate{x, y}] = i
	g.minX = Min(g.minX, x)
	g.minY = Min(g.minY, y)
	g.maxX = Max(g.maxX, x)
	g.maxY = Max(g.maxY, y)
}

// Get retrieves the value located at (x, y)
func (g Grid) Get(x, y int) int {
	return g.field[Coordinate{x, y}]
}

// Len returns the total number of locations stored in the grid
func (g Grid) Len() int {
	return len(g.field)
}

// Print displays the entire grid to STDOUT
func (g Grid) Print() {
	output := [][]string{}
	h, w := g.maxY-g.minY, g.maxX-g.minX

	for i := 0; i <= h; i++ {
		row := []string{}
		for j := 0; j <= w; j++ {
			row = append(row, " ")
		}
		output = append(output, row)
	}

	for l, c := range g.field {
		output[l.Y+Abs(g.minY)][l.X+Abs(g.minX)] = strconv.Itoa(c)
	}

	for i := range output {
		for j := range output[i] {
			fmt.Print(output[i][j])
		}
		fmt.Print("\n")
	}
}
