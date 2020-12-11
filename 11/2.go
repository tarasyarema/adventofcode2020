package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const file = "data.txt"

func main() {
	fileName := file

	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	maze := make([][]rune, 0)
	next := make([][]rune, 0)

	for i, rawLine := range strings.Split(string(b), "\n") {
		line := strings.TrimSpace(rawLine)

		if line == "" {
			continue
		}

		maze = append(maze, make([]rune, len(line)))
		next = append(next, make([]rune, len(line)))
		for j, c := range line {
			maze[i][j] = c
		}
	}

	count := 0
	its := 0

	for {
		change := false
		count = 0

		for i := range maze {
			for j := range maze[i] {
				switch maze[i][j] {
				case 'L':
					if ok, _ := getAdjacent(maze, i, j); ok {
						change = true
						next[i][j] = '#'
					} else {
						next[i][j] = maze[i][j]
					}
				case '#':
					if _, n := getAdjacent(maze, i, j); n > 4 {
						change = true
						next[i][j] = 'L'
					} else {
						next[i][j] = maze[i][j]
					}
				default:
					next[i][j] = maze[i][j]
				}

				if next[i][j] == '#' {
					count++
				}
			}
		}

		its++

		if !change {
			break
		}

		maze, next = next, maze
	}

	fmt.Printf("%d its\n", its)
	fmt.Println(count)
}

func printMaze(maze [][]rune) {
	for i := range maze {
		for _, c := range maze[i] {
			fmt.Printf("%c", c)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func getAdjacent(m [][]rune, i, j int) (bool, int) {
	adjacent := true
	count := 0

	out := false

	// Horizontal W
	for k := j - 1; k >= 0; k-- {
		switch m[i][k] {
		case '#':
			adjacent = false
			count++
			out = true
		case 'L':
			out = true
		}

		if out {
			break
		}
	}

	// Horizontal E
	out = false
	for k := j + 1; k < len(m[i]); k++ {
		switch m[i][k] {
		case '#':
			adjacent = false
			count++
			out = true
		case 'L':
			out = true
		}

		if out {
			break
		}
	}

	// Vertical N
	out = false
	for k := i - 1; k >= 0; k-- {
		switch m[k][j] {
		case '#':
			adjacent = false
			count++
			out = true
		case 'L':
			out = true
		}

		if out {
			break
		}
	}

	// Vertical S
	out = false
	for k := i + 1; k < len(m); k++ {
		switch m[k][j] {
		case '#':
			adjacent = false
			count++
			out = true
		case 'L':
			out = true
		}

		if out {
			break
		}
	}

	// Digonal NW
	l := 1
	out = false
	for k := j - 1; k >= 0 && i-l >= 0; k-- {
		switch m[i-l][k] {
		case '#':
			adjacent = false
			count++
			out = true
		case 'L':
			out = true
		}

		if out {
			break
		}
		l++
	}

	// Diagonal SE
	l = 1
	out = false
	for k := j + 1; k < len(m[i]) && i+l < len(m); k++ {
		switch m[i+l][k] {
		case '#':
			adjacent = false
			count++
			out = true
		case 'L':
			out = true
		}

		if out {
			break
		}
		l++
	}

	// Diagonal SW
	l = 1
	out = false
	for k := i + 1; k < len(m) && j-l >= 0; k++ {
		switch m[k][j-l] {
		case '#':
			adjacent = false
			count++
			out = true
		case 'L':
			out = true
		}

		if out {
			break
		}
		l++
	}

	// Diagonal NE
	l = 1
	out = false
	for k := i - 1; k >= 0 && j+l < len(m[i]); k-- {
		switch m[k][j+l] {
		case '#':
			adjacent = false
			count++
			out = true
		case 'L':
			out = true
		}

		if out {
			break
		}
		l++
	}

	return adjacent, count
}
