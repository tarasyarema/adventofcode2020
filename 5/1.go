package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const file = "data1.txt"

type id struct {
	seat, id int
}

func getID(s string) id {
	var (
		x    = [2]int{0, 127}
		y    = [2]int{0, 7}
		seat int
	)

	for _, c := range s[:7] {
		mid := (x[1] - x[0]) / 2
		switch c {
		case 'F':
			x[1] = x[0] + mid
		case 'B':
			x[0] = x[1] - mid
		}
	}

	for _, c := range s[7:] {
		mid := (y[1] - y[0]) / 2
		switch c {
		case 'L':
			y[1] = y[0] + mid
		case 'R':
			y[0] = y[1] - mid
		}
	}

	seat = x[0] + y[0]
	return id{seat, x[0]*8 + y[0]}
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

	max := 0
	ids := make([]bool, 127*8+7)

	for _, line := range strings.Split(string(b), "\n") {
		if line == "" {
			continue
		}

		current := getID(line)

		if current.id > max {
			max = current.id
		}

		ids[current.id] = true
	}

	fmt.Println("1", max)

	naiveOffset := 100
	for j := range ids[naiveOffset:] {
		i := naiveOffset + j
		if !ids[i] {
			fmt.Println("2", i)
			break
		}
	}

}
