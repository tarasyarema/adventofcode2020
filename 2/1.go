package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const file = "data1.txt"

func main() {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	ok := 0

	for _, x := range strings.Split(string(b), "\n") {
		if x == "" {
			continue
		}

		line := strings.Split(x, ": ")
		tmp := strings.Split(line[0], " ")

		var (
			min, max int64
		)

		for i, n := range strings.Split(tmp[0], "-") {
			if i == 0 {
				min, _ = strconv.ParseInt(n, 0, 64)
			} else {
				max, _ = strconv.ParseInt(n, 0, 64)
			}
		}

		letter := rune(tmp[1][0])
		password := line[1]

		var current int64 = 0

		for _, r := range password {
			if r == letter {
				current += 1
			}
		}

		if current >= min && current <= max {
			ok += 1
		}
	}

	fmt.Println(ok)
}
