package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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

	t := 0
	times := make([]int, 0)

	for i, rawLine := range strings.Split(string(b), "\n") {
		line := strings.TrimSpace(rawLine)

		if line == "" {
			continue
		}

		if i == 0 {
			n, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			t = n
		} else {
			for _, b := range strings.Split(line, ",") {
				if b != "x" {
					n, err := strconv.Atoi(b)
					if err != nil {
						panic(err)
					}

					times = append(times, n)
				}
			}
		}
	}

	i := t
	sol := 0

	for {
		found := false
		for _, tt := range times {
			if i%tt == 0 {
				sol = (i - t) * tt
				found = true
				break
			}
		}

		if found {
			break
		}

		i++
	}

	fmt.Println(sol)
}
