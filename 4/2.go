package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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

	n := 1
	count := 0
	type field struct {
		re   string
		have bool
	}

	fields := []field{
		{`byr:(19[2-9][0-9]|2000|2001|2002)`, false},
		{`iyr:20(1[0-9]|20)`, false},
		{`eyr:20(2[0-9]|30)`, false},
		{`hgt:(1([5-8][0-9]|9[0-3])cm|(59|6[0-9]|7[0-6])in)`, false},
		{`hcl:#[0-9a-f]{6}`, false},
		{`ecl:(amb|blu|brn|gry|grn|hzl|oth)`, false},
		{`pid:[0-9]{9}[^0-9a-zA-Z]`, false},
	}

	for _, line := range strings.Split(string(b), "\n") {
		if len(line) < 2 {
			bad := false
			for _, f := range fields {
				if !f.have {
					bad = true
					break
				}
			}

			if !bad {
				count += 1
			}

			for p := range fields {
				fields[p].have = false
			}

			n += 1
			continue
		}

		for p, f := range fields {
			if ok, err := regexp.MatchString(f.re, line); err != nil {
				panic(err)
			} else {
				if ok {
					fields[p].have = true
				}
			}
		}
	}

	fmt.Println(count)
}
