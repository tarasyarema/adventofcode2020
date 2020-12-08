package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

const file = "data1.txt"

type Op struct {
	T       string
	N       int
	Ran     bool
	Swapped bool
}

type Program struct {
	Running bool
	Pos     int
	Ops     []Op
	Acc     int
	Swapped bool
	Looped  bool
}

func NewProgramFromBytes(b []byte) *Program {
	ops := make([]Op, 0)

	for _, rawLine := range strings.Split(string(b), "\n") {
		line := strings.TrimSpace(rawLine)
		if line == "" {
			continue
		}

		data := strings.Split(line, " ")

		n, err := strconv.Atoi(data[1])
		if err != nil {
			panic(err)
		}

		ops = append(ops, Op{N: n, T: data[0], Ran: false, Swapped: false})
	}

	return &Program{
		Running: true,
		Pos:     0,
		Ops:     ops,
		Acc:     0,
		Swapped: false,
		Looped:  false,
	}
}

func (p *Program) copyProgram() *Program {
	p2 := &Program{
		Running: p.Running,
		Pos:     p.Pos,
		Ops:     make([]Op, len(p.Ops)),
		Acc:     p.Acc,
		Swapped: p.Swapped,
		Looped:  p.Looped,
	}

	copy(p2.Ops, p.Ops)
	return p2
}

// Next runs the program in the infinite loop way (Part 1)
func (p *Program) Next() bool {
	if !p.Running {
		return false
	}

	if p.Ops[p.Pos].Ran {
		p.Looped = true
		p.Running = false
		return false
	}

	// Mark the current position as ran
	p.Ops[p.Pos].Ran = true

	switch p.Ops[p.Pos].T {
	case "nop":
		p.Pos++
	case "acc":
		p.Acc += p.Ops[p.Pos].N
		p.Pos++
	case "jmp":
		p.Pos += p.Ops[p.Pos].N
	}

	return true
}

// Fix is the wrapper function to handle new swap branches
func (p *Program) Fix() {
	defer wg.Done()

	for {
		if !p.NextBranch() {
			break
		}
	}

	// If the current routine has ended and
	// there was no loops, we know we fixed the program
	if !p.Looped {
		fmt.Println("2", p.Acc)
	}
}

// NextBranch is the a function that works like Next but
// it creates sub-goruotines for every swapping branch of ops
func (p *Program) NextBranch() bool {
	// We check for the "fixed" solution
	// as the program is expected to end normally
	// and we mark it as not looped
	if p.Pos >= len(p.Ops) {
		p.Looped = false
		p.Running = false
	}

	if !p.Running {
		return false
	}

	// If the program is corrupted we will end up in a loop
	// so this check for this condition and marks the program as looped
	if p.Ops[p.Pos].Ran {
		p.Running = false
		p.Looped = true
		return false
	}

	// Mark the current position as ran
	p.Ops[p.Pos].Ran = true

	switch p.Ops[p.Pos].T {
	case "nop":
		// If the current program (goroutine) had not swaps
		// we swap the op and create a new goroutine
		if !p.Swapped {
			branch := p.copyProgram()

			branch.Ops[p.Pos].T = "jmp"
			p.Ops[p.Pos].Ran = false
			branch.Swapped = true

			wg.Add(1)
			go branch.Fix()
		}

		p.Pos++
	case "acc":
		p.Acc += p.Ops[p.Pos].N
		p.Pos++
	case "jmp":
		// Same as for the "nop" case
		if !p.Swapped {
			branch := p.copyProgram()

			branch.Ops[p.Pos].T = "nop"
			branch.Ops[p.Pos].Ran = false
			branch.Swapped = true

			wg.Add(1)
			go branch.Fix()
		}

		p.Pos += p.Ops[p.Pos].N
	}

	return true
}

func main() {
	fileName := file

	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	p := NewProgramFromBytes(b)

	for {
		if !p.Next() {
			fmt.Println("1", p.Acc)
			break
		}
	}

	p = NewProgramFromBytes(b)

	// Create the first program
	wg.Add(1)
	go p.Fix()

	// Wait for all the swap goruotines to end
	wg.Wait()
}
