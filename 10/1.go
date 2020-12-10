package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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

	adapters := make([]int, 0)

	for _, rawLine := range strings.Split(string(b), "\n") {
		line := strings.TrimSpace(rawLine)

		if line == "" {
			continue
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		adapters = append(adapters, n)
	}

	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	begin := 0
	singular := 0
	trio := 0

	for i := 0; i < len(adapters); i++ {
		delta := adapters[i] - begin
		switch delta {
		case 1:
			singular++
		case 3:
			trio++
		}
		begin += delta
	}

	fmt.Println(1, singular*trio)

	adapters = append([]int{0}, adapters...)
	dp := make([]int, len(adapters))

	dp[0] = 1
	dp[1] = 1

	for i := 2; i < len(dp); i++ {
		for j := i - 1; j >= 0; j-- {
			if adapters[i]-adapters[j] <= 3 {
				dp[i] += dp[j]
			} else {
				break
			}
		}
	}

	fmt.Println(2, dp[len(dp)-1])
}

func recurse(a []int) int {
	sum := 1

	for i := 1; i < len(a); i++ {
		if a[i]-a[0] > 3 {
			break
		}

		sum += recurse(a[i:])
	}

	return sum
}
