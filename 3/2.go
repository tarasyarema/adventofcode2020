package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const file = "data1.txt"

type position struct {
	row, col, rowInc, colInc, found int
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

	positions := []position{
		{0, 0, 1, 1, 0},
		{0, 0, 3, 1, 0},
		{0, 0, 5, 1, 0},
		{0, 0, 7, 1, 0},
		{0, 0, 1, 2, 0},
	}

	for i, line := range strings.Split(string(b), "\n") {
		if line == "" {
			continue
		}

		for p := range positions {
			//Check if we even need to check
			if i != positions[p].col {
				continue
			}

			// Check for tree
			if line[positions[p].row] == '#' {
				positions[p].found += 1
			}

			// Make next position update
			positions[p].row = (positions[p].row + positions[p].rowInc) % (len(line) - 1)
			positions[p].col += positions[p].colInc
		}
	}

	result := 1
	for _, position := range positions {
		result *= position.found
	}

	fmt.Println(result)
}
