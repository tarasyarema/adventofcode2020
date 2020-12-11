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
					if getAdjacentL(maze, i, j) {
						change = true
						next[i][j] = '#'
					} else {
						next[i][j] = maze[i][j]
					}
				case '#':
					if getAdjacentH(maze, i, j) {
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

func getAdjacentL(m [][]rune, i, j int) bool {
	adjacent := 1

	for v := i - 1; v <= i+1; v++ {
		if v < 0 || v >= len(m) {
			continue
		}

		for h := j - 1; h <= j+1; h++ {
			if h < 0 || h >= len(m[i]) || (v == i && h == j) {
				continue
			}

			if m[v][h] == '#' {
				adjacent *= 0
			}
		}
	}

	if adjacent == 1 {
		return true
	}

	return false
}

func getAdjacentH(m [][]rune, i, j int) bool {
	adjacent := 0

	for v := i - 1; v <= i+1; v++ {
		if v < 0 || v >= len(m) {
			continue
		}

		for h := j - 1; h <= j+1; h++ {
			if h < 0 || h >= len(m[i]) || (v == i && h == j) {
				continue
			}

			if m[v][h] == '#' {
				adjacent += 1
			}
		}
	}

	if adjacent >= 4 {
		return true
	}

	return false
}
