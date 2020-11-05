package aoc

import (
	"log"
	"math"
)

// Program represents an initialized Intcode program.
type Program struct {
	init, state map[int]int
	pc, op, ro  int
	Input       <-chan int
	Output      chan<- int
}

type operation struct {
	code  int
	modes [3]bool
}

// NewProgram accepts an initial state and returns a ready-to-run Program, as well as input and output channels for the caller to utilize.
func NewProgram(i []int) (*Program, chan<- int, <-chan int) {
	init, state := map[int]int{}, map[int]int{}
	in, out := newInChan(), newOutChan()

	for i, v := range i {
		state[i] = v
		init[i] = v
	}

	p := Program{
		init:   init,
		state:  state,
		Input:  in,
		Output: out,
	}

	return &p, in, out
}

// Reset returns the program to its initial state.  Since a correctly exited program stops by closing the output channel, and since unread inputs may exist, this method creates and returns new input/output channels.
func (p *Program) Reset() (chan<- int, <-chan int) {
	state := map[int]int{}
	for i, v := range p.init {
		state[i] = v
	}

	p.pc, p.op, p.ro = 0, 0, 0

	in, out := newInChan(), newOutChan()
	p.Input = in
	p.Output = out

	return in, out
}

// Set stores integer i at state address n
func (p *Program) Set(n, i int) {
	p.state[n] = 2
}

func newInChan() chan int {
	return make(chan int, 1)
}

func newOutChan() chan int {
	return make(chan int, 1)
}

// Run executes an Intcode program until a halt operation is encountered.
func (p *Program) Run() {
	var done bool
	for !done {
		done = p.Step()
	}
}

// Step executes a single instruction and returns true if the halt signal was encountered.
func (p *Program) Step() bool {
	i := p.state[p.pc]
	op := i % 100

	switch op {
	case 1:
		// Add
		a, b, c := p.getParam(1), p.getParam(2), p.state[p.pc+3]
		p.putParam(3, c, a+b)
		p.pc += 4
	case 2:
		// Multiply
		a, b, c := p.getParam(1), p.getParam(2), p.state[p.pc+3]
		p.putParam(3, c, a*b)
		p.pc += 4
	case 3:
		// Input
		a := p.state[p.pc+1]
		p.putParam(1, a, <-p.Input)
		p.pc += 2
	case 4:
		// Output
		a := p.getParam(1)
		p.Output <- a
		p.pc += 2
	case 5:
		// Jump if true
		a, b := p.getParam(1), p.getParam(2)
		if a != 0 {
			p.pc = b
		} else {
			p.pc += 3
		}
	case 6:
		// Jump if false
		a, b := p.getParam(1), p.getParam(2)
		if a == 0 {
			p.pc = b
		} else {
			p.pc += 3
		}
	case 7:
		// Less than
		a, b, c := p.getParam(1), p.getParam(2), p.state[p.pc+3]
		if a < b {
			p.putParam(3, c, 1)
		} else {
			p.putParam(3, c, 0)
		}
		p.pc += 4
	case 8:
		// Equals
		a, b, c := p.getParam(1), p.getParam(2), p.state[p.pc+3]
		if a == b {
			p.putParam(3, c, 1)
		} else {
			p.putParam(3, c, 0)
		}
		p.pc += 4
	case 9:
		// Relative offset
		a := p.getParam(1)
		p.ro += a
		p.pc += 2
	case 99:
		// Halt
		close(p.Output)
		return true
	default:
		log.Fatalf("encountered unknown opcode; %+v", p)
	}

	return false
}

func (p *Program) getParam(o int) int {
	param := p.state[p.pc+o]
	mode := (p.state[p.pc] / int(math.Pow(float64(10), float64(o+1)))) % 10

	switch mode {
	case 0:
		return p.state[param]
	case 1:
		return param
	case 2:
		return p.state[p.ro+param]
	default:
		log.Fatalf("Encountered unknown mode; %+v", p)
		return -1
	}
}

func (p *Program) putParam(pos, adr, val int) {
	mode := (p.state[p.pc] / int(math.Pow(float64(10), float64(pos+1)))) % 10

	switch mode {
	case 0:
		p.state[adr] = val
	case 2:
		p.state[p.ro+adr] = val
	default:
		log.Fatalf("Encountered unknown mode; %+v", p)
	}
}
