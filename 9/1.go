package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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

	stackLen := 25
	stack := make([]int64, stackLen)

	for i, rawLine := range strings.Split(string(b), "\n") {
		line := strings.TrimSpace(rawLine)

		if line == "" {
			continue
		}

		n, err := strconv.ParseInt(line, 0, 64)
		if err != nil {
			panic(err)
		}

		if i < stackLen {
			stack[i] = n
			continue
		}

		ok := false

		for j := i - stackLen; j < len(stack); j++ {
			if ok {
				break
			}

			for k := j + 1; k < len(stack); k++ {
				if j != k {
					if stack[j]+stack[k] == n {
						ok = true
						break
					}
				}
			}
		}

		if !ok {
			fmt.Println("1 ->", n)

			r := [2]int{0, 0}

			for j := 0; j < len(stack); j++ {
				if r[0] != 0 || r[1] != 0 {
					break
				}

				s := int64(0)

				for k := j + 1; k < len(stack); k++ {
					s += stack[k]

					if s == n {
						r[0] = j
						r[1] = k
						break
					}
				}
			}

			max := int64(0)
			min := int64(9223372036854775807)

			for j := r[0] + 1; j <= r[1]; j++ {
				if stack[j] > max {
					max = stack[j]
				}
				if stack[j] < min {
					min = stack[j]
				}
			}

			fmt.Println("2 ->", max+min)
			break
		}

		stack = append(stack, n)
	}
}
