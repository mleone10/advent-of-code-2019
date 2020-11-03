package aoc

import (
	"log"
	"math"
)

// Program represents an initialized, running Intcode program
type Program struct {
	state      map[int]int
	pc, op, ro int
	Input      <-chan int
	Output     chan<- int
}

type operation struct {
	code  int
	modes [3]bool
}

// NewProgram constructs a Program with the given initial state and input, output, and halt channels
func NewProgram(init []int, in <-chan int, out chan<- int) *Program {
	state := map[int]int{}
	for i, v := range init {
		state[i] = v
	}

	return &Program{
		state:  state,
		Input:  in,
		Output: out,
	}
}

// Run executes an Intcode program until a halt operation is encountered
func (p *Program) Run() {
	for {
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
			return
		default:
			log.Fatalf("encountered unknown opcode; %+v", p)
		}
	}
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
