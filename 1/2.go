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

	data := make([]int64, 0)
	for i, x := range strings.Split(string(b), "\n") {
		if x == "" {
			continue
		}

		n, err := strconv.ParseInt(x, 0, 0)
		if err != nil {
			fmt.Printf("%d (%s): %s\n", i, x, err.Error())
			panic(err)
		}

		data = append(data, n)
	}

	done := false
	want := []int64{0, 0, 0}

	for _, x := range data {
		if done {
			break
		}
		for _, y := range data {
			for _, z := range data {
				if x+y+z == 2020 {
					want = []int64{x, y, z}
					done = true
					break
				}
			}
		}
	}

	fmt.Println(want[0] * want[1] * want[2])
}
