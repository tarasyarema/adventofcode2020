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

	n := 1
	ok := 0
	fields := map[string]bool{
		"byr": false,
		"iyr": false,
		"eyr": false,
		"hgt": false,
		"hcl": false,
		"ecl": false,
		"pid": false,
		// "cid": false,
	}

	for _, line := range strings.Split(string(b), "\n") {
		if len(line) < 2 {
			bad := false
			for _, is := range fields {
				if !is {
					bad = true
					break
				}
			}

			if !bad {
				ok += 1
			}

			for field := range fields {
				fields[field] = false
			}

			n += 1
			continue
		}

		for field := range fields {
			if strings.Contains(line, field+":") {
				fields[field] = true
			}
		}
	}

	fmt.Println(ok)
}
