package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const file = "data.txt"

func getMask(m []int) uint64 {
	if len(m) != 36 {
		panic("bad len")
	}

	mask := uint64(0)

	for i := len(m) - 1; i >= 0; i-- {
		b := m[i]
		if b == 1 {
			mask += uint64(math.Pow(2, float64(len(m)-i-1)))
		}
	}

	return mask
}

func computeMask(s string) []int {
	if len(s) != 36 {
		panic("wrong len")
	}

	m := make([]int, 36)

	for i := range s {
		switch s[i] {
		case '1':
			m[i] = 1
		case '0':
			m[i] = 0
		}
	}

	return m
}

func getUint(v []int) uint64 {
	if len(v) != 36 {
		panic("wrong len")
	}
	r := uint64(0)

	for i := 0; i < len(v); i++ {
		if v[i] == 0 {
			continue
		}
		r += uint64(math.Floor(math.Pow(float64(2), float64(len(v)-i-1))))
	}

	return r
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

	mask := make([]int, 36)
	mem := map[int]uint64{}

	maskRe := regexp.MustCompile(`mask = (?P<mask>[01X]{36})`)
	memRe := regexp.MustCompile(`mem\[(?P<addr>[0-9]+)\] = (?P<val>[0-9]+)`)

	for _, rawLine := range strings.Split(string(b), "\n") {
		line := strings.TrimSpace(rawLine)

		if line == "" {
			continue
		}

		if ok := maskRe.MatchString(line); ok {
			m := maskRe.FindStringSubmatch(line)[1]

			fmt.Println("mask", m)

			for i, c := range m {
				switch c {
				case 'X':
					mask[i] = 2
				case '1':
					mask[i] = 1
				case '0':
					mask[i] = 0
				}
			}

			fmt.Println(mask)
		}

		if ok := memRe.MatchString(line); ok {
			matches := memRe.FindStringSubmatch(line)

			addrString := matches[1]
			valString := matches[2]

			addr, err := strconv.Atoi(addrString)
			if err != nil {
				panic(err)
			}

			val, err := strconv.ParseUint(valString, 0, 64)
			if err != nil {
				panic(err)
			}

			x := computeMask(fmt.Sprintf("%036b", val))
			res := make([]int, 36)

			for i := 0; i < 36; i++ {
				res[i] = x[i]

				switch mask[i] {
				case 1:
					res[i] = 1
				case 0:
					res[i] = 0
				}
			}

			fmt.Println("mem", addr, val, getUint(res))
			mem[addr] = getUint(res)
		}
	}

	fmt.Println(mem)

	sol := uint64(0)

	for i := range mem {
		sol += mem[i]
	}

	fmt.Println(sol)
}

func apply(value uint64, mask []int) {
	s := ""
	for i := 0; i < len(mask); i++ {
		b := value & (1 << uint(i))

		switch b {
		case 0:
			s = fmt.Sprintf("%d%s", mask[len(mask)-i-1], s)
		case 1:
			s = fmt.Sprintf("%d%s", b, s)
		}

	}

	fmt.Println(s)
}

func printBits(x uint64) {
	fmt.Printf("%036b = %d\n", x, x)
}
