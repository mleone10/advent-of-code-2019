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

	p, in, out := aoc.NewProgram(init)
	go p.Run()

	grid := interact(in, out, 0)
	log.Printf("Total panels painted at least once: %d", grid.Len())

	in, out = p.Reset()
	go p.Run()
	interact(in, out, 1).Print()
}

func interact(in chan<- int, out <-chan int, initColor int) aoc.Grid {
	var grid aoc.Grid
	var dir direction
	var x, y, numOuts, paintColor int

	grid.Mapper = mappingFunc
	grid.Set(0, 0, initColor)
	in <- initColor

	for {
		select {
		case o, ok := <-out:
			if !ok {
				return grid
			}
			numOuts++
			if numOuts == 1 {
				grid.Set(x, y, o)
			} else if numOuts == 2 {
				numOuts = 0
				if o == 0 {
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
				}

				if grid.Get(x, y) == 0 {
					paintColor = 0
				} else {
					paintColor = 1
				}
			}
		case in <- paintColor:
		}
	}

	// for color := range out {
	// 	grid.Set(x, y, color)

	// 	rotation := <-out
	// 	if rotation == 0 {
	// 		dir = (dir + 3) % 4
	// 	} else {
	// 		dir = (dir + 5) % 4
	// 	}

	// 	switch dir {
	// 	case dirUp:
	// 		y--
	// 	case dirRight:
	// 		x++
	// 	case dirDown:
	// 		y++
	// 	case dirLeft:
	// 		x--
	// 	default:
	// 		log.Fatalf("Encountered unsupported direction %d", dir)
	// 	}

	// 	if grid.Get(x, y) == 0 {
	// 		in <- 0
	// 	} else {
	// 		in <- 1
	// 	}
	// }
}

func mappingFunc(i int) string {
	if i == 1 {
		return "#"
	}
	return " "
}
