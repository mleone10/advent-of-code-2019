package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	aoc "github.com/mleone10/advent-of-code-2019"
)

type direction int
type status int

const (
	dirNorth direction = 1
	dirSouth direction = 2
	dirWest  direction = 3
	dirEast  direction = 4
)

const (
	statWall  status = 0
	statMoved status = 1
	statOxy   status = 2
)

type droid struct {
	in  chan<- int
	out <-chan int
}

type node struct {
	l aoc.Coordinate
	n int
}

var mVectors = map[direction]aoc.Coordinate{
	dirNorth: aoc.Coordinate{X: 0, Y: -1},
	dirSouth: aoc.Coordinate{X: 0, Y: 1},
	dirWest:  aoc.Coordinate{X: -1, Y: 0},
	dirEast:  aoc.Coordinate{X: 1, Y: 0},
}

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

	d := droid{in, out}
	var visited aoc.Grid
	d.visit(aoc.Coordinate{}, &visited)
	q := findOxygen(visited)

	log.Printf("Shortest distance from droid start position to oxygen system: %d", findShortestPath(aoc.Coordinate{}, q, visited))
	log.Printf("Time to fill entire area: %d", findTimeToFill(findOxygen(visited), visited))
}

func (d droid) step(dir direction) status {
	d.in <- int(dir)
	return status(<-d.out)
}

func (d droid) visit(loc aoc.Coordinate, visited *aoc.Grid) {
	// Recursively explore all unvisited surrounding locations
	for dir, v := range mVectors {
		nextLoc := loc.Add(v)
		if _, ok := visited.Field[nextLoc]; ok {
			// Already visited nextLoc, move on
			continue
		}
		stat := d.step(dir)
		if visited.SetCoord(nextLoc, int(stat)); stat == statWall {
			// NextLoc is a wall, can't visit there
			continue
		}
		d.visit(nextLoc, visited)
		// We've visited all new locations around this one, move back to where we were
		d.step(opposite(dir))
	}
}

func calculateDistancesFrom(origin aoc.Coordinate, maze aoc.Grid) map[aoc.Coordinate]int {
	visited := map[aoc.Coordinate]bool{}
	dist := map[aoc.Coordinate]int{}

	searchQueue := []node{{origin, 0}}
	var n node

	for len(searchQueue) > 0 {
		n, searchQueue = searchQueue[0], searchQueue[1:]
		dist[n.l] = n.n

		for _, v := range mVectors {
			neighbor := n.l.Add(v)
			if visited[neighbor] {
				continue
			}
			visited[neighbor] = true
			i, _ := maze.Field[neighbor]
			if i != int(statWall) {
				searchQueue = append(searchQueue, node{neighbor, n.n + 1})
			}
		}
	}

	return dist
}

func findOxygen(maze aoc.Grid) aoc.Coordinate {
	for l, v := range maze.Field {
		if v == int(statOxy) {
			return l
		}
	}
	log.Fatalln("Could not found oxygen in maze")
	return aoc.Coordinate{}
}

func findShortestPath(from, to aoc.Coordinate, maze aoc.Grid) int {
	return calculateDistancesFrom(from, maze)[to]
}

func findTimeToFill(from aoc.Coordinate, maze aoc.Grid) int {
	max := 0
	for _, t := range calculateDistancesFrom(from, maze) {
		max = aoc.Max(max, t)
	}
	return max
}

func opposite(dir direction) direction {
	switch dir {
	case dirNorth:
		return dirSouth
	case dirSouth:
		return dirNorth
	case dirEast:
		return dirWest
	case dirWest:
		return dirEast
	default:
		log.Fatalf("Encountered invalid direction %d", dir)
		return 0
	}
}
