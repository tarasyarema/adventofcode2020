package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const file = "data1.txt"

func main() {
	fileName := file

	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	position := 0
	found := 0

	for _, line := range strings.Split(string(b), "\n") {
		if line == "" {
			continue
		}

		if line[position] == '#' {
			found += 1
		}

		for i := range line {
			if i == position {
				fmt.Printf("O")
			} else {
				fmt.Printf("%s", string(line[i]))
			}
		}
		fmt.Printf("\n")

		position = (position + 3) % (len(line) - 1)
	}

	fmt.Println(found)
}
