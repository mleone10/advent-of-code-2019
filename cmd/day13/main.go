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
	var x, y, t, score int
	var grid aoc.Grid
	var tileLoc, ballLoc aoc.Coordinate

	for x = range out {
		y, t = <-out, <-out

		if x == -1 && y == 0 && t > 4 {
			score = t
		} else {
			grid.Set(x, y, t)
		}

		switch t {
		case tilePaddle:
			tileLoc.X, tileLoc.Y = x, y
		case tileBlock:
			ballLoc.X, ballLoc.Y = x, y
		}

		if ballLoc.Y > tileLoc.Y {
			in <- moveRight
		} else if ballLoc.Y < tileLoc.Y {
			in <- moveLeft
		} else {
			in <- moveNeutral
		}
	}

	return score
}
