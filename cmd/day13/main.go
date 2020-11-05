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
	tileEmpty = iota
	tileWall
	tileBlock
	tilePaddle
	tileBall
)

const (
	moveLeft = iota - 1
	moveNeutral
	moveRight
)

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
	log.Printf("Blocks on screen at start of game: %d", getInitialBlocks(in, out))

	in, out = p.Reset()
	p.Set(0, 2)
	go p.Run()
	log.Printf("Score at end of game: %d", playGame(in, out))
}

func getInitialBlocks(in chan<- int, out <-chan int) int {
	var count int
	for range out {
		<-out
		if <-out == tileBlock {
			count++
		}
	}
	return count
}

func playGame(in chan<- int, out <-chan int) int {
	var x, y, t, move, numOuts, score int
	var grid aoc.Grid
	var paddleLoc, ballLoc aoc.Coordinate

	for {
		select {
		case o, ok := <-out:
			if !ok {
				return score
			}
			numOuts++
			if numOuts == 1 {
				x = o
			} else if numOuts == 2 {
				y = o
			} else if numOuts == 3 {
				t = o
				numOuts = 0

				if x == -1 && y == 0 {
					score = t
				} else {
					grid.Set(x, y, t)
				}

				switch t {
				case tilePaddle:
					paddleLoc.X, paddleLoc.Y = x, y
				case tileBall:
					ballLoc.X, ballLoc.Y = x, y
				}

				if ballLoc.X > paddleLoc.X {
					move = moveRight
				} else if ballLoc.X < paddleLoc.X {
					move = moveLeft
				} else {
					move = moveNeutral
				}
			}
		case in <- move:
		}
	}
}
