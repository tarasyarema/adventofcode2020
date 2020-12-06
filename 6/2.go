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

	sol := 0

	count := 0
	ans := make([]int, 'z'-'a'+1)

	for _, rawLine := range strings.Split(string(b), "\n") {
		line := strings.TrimSpace(rawLine)
		if line == "" {
			for i := range ans {
				if ans[i] == count {
					sol += 1
				}

				ans[i] = 0
			}

			count = 0

			continue
		}

		for _, c := range line {
			if c >= 97 {
				ans[c-97] += 1
			}
		}

		count += 1
	}

	fmt.Println(sol)
}
