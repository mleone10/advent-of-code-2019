package aoc

import (
	"fmt"
	"strconv"
)

// MappingFunc represents a function used to determine which character Print should use for a given integer in the grid.
type MappingFunc func(int) string

// Grid represents a two dimensional array of integers.
type Grid struct {
	// Mapper is a MappingFunc used by Print to convert grid values into more readable characters.  It defaults to just printing the integer.
	Mapper                 MappingFunc
	field                  map[Coordinate]int
	minX, minY, maxX, maxY int
}

// Coordinate represents a two-dimensional (x,y) position on the grid.
type Coordinate struct {
	X, Y int
}

func defaultMappingFunc(i int) string {
	return strconv.Itoa(i)
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

// SetCoord is a convenience method which uses a given Coordinate's (x, y) location to call Set
func (g *Grid) SetCoord(c Coordinate, i int) {
	g.Set(c.X, c.Y, i)
}

// Get retrieves the value located at (x, y)
func (g Grid) Get(x, y int) int {
	return g.field[Coordinate{x, y}]
}

// GetCoord is a convenience method which uses a given Coordinate's (x, y) location to call Get
func (g Grid) GetCoord(c Coordinate) int {
	return g.Get(c.X, c.Y)
}

// Len returns the total number of locations stored in the grid
func (g Grid) Len() int {
	return len(g.field)
}

// Print displays the entire grid to STDOUT using the grid's designated MapperFunc
func (g Grid) Print() {
	if g.Mapper == nil {
		g.Mapper = defaultMappingFunc
	}

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
		output[l.Y+Abs(g.minY)][l.X+Abs(g.minX)] = g.Mapper(c)
	}

	for i := range output {
		for j := range output[i] {
			fmt.Print(output[i][j])
		}
		fmt.Print("\n")
	}
}

// PlusCoord adds two Coordinates together and returns the result.
func PlusCoord(r, s Coordinate) Coordinate {
	return Coordinate{
		X: r.X + s.X,
		Y: r.Y + s.Y,
	}
}

// MinusCoord subtracts Coordinate r from s and returns the result.
func MinusCoord(r, s Coordinate) Coordinate {
	return Coordinate{
		X: r.X - s.X,
		Y: r.Y - s.Y,
	}
}
