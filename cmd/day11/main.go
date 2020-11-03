package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	aoc "github.com/mleone10/advent-of-code-2019"
)

const (
	dirUp = iota
	dirRight
	dirDown
	dirLeft
)

type location struct {
	x, y int
}

type grid struct {
	ls               map[location]int
	w, h, minX, minY int
}

type direction int

func main() {
	var init []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l := scanner.Text()

	for _, s := range strings.Split(l, ",") {
		i, _ := strconv.Atoi(s)
		init = append(init, i)
	}

	grid := run(init, 0)
	log.Printf("Total panels painted at least once: %d", len(grid.ls))

	displayGrid(run(init, 1))
}

func run(prog []int, initColor int) grid {
	in := make(chan int, 1)
	out := make(chan int)
	p := aoc.NewProgram(prog, in, out)

	go p.Run()
	return interact(in, out, initColor)
}

func interact(in chan<- int, out <-chan int, initColor int) grid {
	grid := grid{map[location]int{}, 1, 1, 0, 0}
	var l location
	var dir direction
	var maxX, maxY int
	minX, minY := math.MaxInt64, math.MaxInt64

	grid.ls[l] = initColor
	in <- initColor

	for color := range out {
		grid.ls[l] = color

		rotation := <-out
		if rotation == 0 {
			dir = (dir + 3) % 4
		} else {
			dir = (dir + 5) % 4
		}

		switch dir {
		case dirUp:
			l.y--
		case dirRight:
			l.x++
		case dirDown:
			l.y++
		case dirLeft:
			l.x--
		default:
			log.Fatalf("Encountered unsupported direction %d", dir)
		}

		minX = aoc.Min(minX, l.x)
		minY = aoc.Min(minY, l.y)
		maxX = aoc.Max(maxX, l.x)
		maxY = aoc.Max(maxY, l.y)

		if grid.ls[l] == 0 {
			in <- 0
		} else {
			in <- 1
		}
	}

	grid.h = aoc.Abs(maxY - minY)
	grid.w = aoc.Abs(maxX - minX)
	grid.minX = minX
	grid.minY = minY

	return grid
}

func displayGrid(g grid) {
	output := [][]string{}

	for i := 0; i <= g.h; i++ {
		row := []string{}
		for j := 0; j <= g.w; j++ {
			row = append(row, " ")
		}
		output = append(output, row)
	}

	for l, c := range g.ls {
		if c == 1 {
			output[l.y+aoc.Abs(g.minY)][l.x+aoc.Abs(g.minX)] = "#"
		}
	}

	for i := range output {
		for j := range output[i] {
			fmt.Print(output[i][j])
		}
		fmt.Print("\n")
	}
}
