package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const file = "data1.txt"

func computeSum(g map[string][]child, current string) int64 {
	sum := int64(1)

	for _, c := range g[current] {
		sum += c.N * computeSum(g, c.C)
	}

	return sum
}

func canHold(g map[string][]child, parent string) (map[string]bool, int64) {
	visited := map[string]bool{}
	visited[parent] = false

	for {
		pivot := ""

		for k, ok := range visited {
			if !ok {
				pivot = k
				break
			}
		}

		if pivot == "" {
			break
		}

		visited[pivot] = true

		for _, neig := range g[pivot] {
			if _, ok := visited[neig.C]; !ok {
				visited[neig.C] = false
			}
		}
	}

	return visited, computeSum(g, parent) - 1
}

type child struct {
	N int64  `json:"n"`
	C string `json:"c"`
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

	sol := int64(0)
	g := map[string][]child{}

	reN := regexp.MustCompile(`[0-9]+`)
	reBag := regexp.MustCompile(`(?P<Color>[a-z]+ [a-z]+) bag`)
	reNoOther := regexp.MustCompile(`no other bags.`)

	for _, rawLine := range strings.Split(string(b), "\n") {
		line := strings.TrimSpace(rawLine)

		if line == "" {
			continue
		}

		parts := strings.Split(line, " bags contain ")

		parent := strings.TrimSpace(parts[0])
		childs := make([]child, 0)

		if reNoOther.MatchString(parts[1]) {
			g[parent] = childs
			continue
		}

		for _, rString := range strings.Split(strings.TrimSpace(parts[1]), ", ") {
			n, err := strconv.ParseInt(reN.FindString(rString), 0, 64)
			if err != nil {
				panic(err)
			}

			bagColor := reBag.FindStringSubmatch(rString)[1]
			childs = append(childs, child{n, bagColor})
		}

		g[parent] = childs
	}

	want := "shiny gold"

	for parent := range g {
		if parent == want {
			continue
		}

		if visited, _ := canHold(g, parent); visited[want] {
			sol += 1
		}
	}

	fmt.Println("1 ->", sol)

	_, sol = canHold(g, want)
	fmt.Println("2 ->", sol)
}
