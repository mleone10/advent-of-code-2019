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

func main() {
	var init []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l := scanner.Text()

	for _, s := range strings.Split(l, ",") {
		i, _ := strconv.Atoi(s)
		init = append(init, i)
	}

	run(init)
}

func run(prog []int) {
	p, in, out := aoc.NewProgram(prog)

	go p.Run()
	log.Printf("Blocks on screen at start of game: %d", getInitialBlocks(in, out))
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
