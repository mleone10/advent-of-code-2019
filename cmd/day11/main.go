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
	dirUp = iota
	dirRight
	dirDown
	dirLeft
)

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
	log.Printf("Total panels painted at least once: %d", grid.Len())

	run(init, 1).Print()
}

func run(prog []int, initColor int) aoc.Grid {
	in := make(chan int, 1)
	out := make(chan int)
	p := aoc.NewProgram(prog, in, out)

	go p.Run()
	return interact(in, out, initColor)
}

func interact(in chan<- int, out <-chan int, initColor int) aoc.Grid {
	var grid aoc.Grid
	var dir direction
	var x, y int

	grid.Set(0, 0, initColor)
	in <- initColor

	for color := range out {
		grid.Set(x, y, color)

		rotation := <-out
		if rotation == 0 {
			dir = (dir + 3) % 4
		} else {
			dir = (dir + 5) % 4
		}

		switch dir {
		case dirUp:
			y--
		case dirRight:
			x++
		case dirDown:
			y++
		case dirLeft:
			x--
		default:
			log.Fatalf("Encountered unsupported direction %d", dir)
		}

		if grid.Get(x, y) == 0 {
			in <- 0
		} else {
			in <- 1
		}
	}

	return grid
}
