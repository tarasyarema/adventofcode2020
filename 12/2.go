package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

const file = "data.txt"

type point struct {
	x, y int
}

type vec struct {
	x, y int
}

type ship struct {
	p *point
	v *vec
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func (s ship) forward(units int) {
	s.p.x += units * s.v.x
	s.p.y += units * s.v.y
}

func (s ship) move(v vec) {
	s.p.x += v.x
	s.p.y += v.y
}

func (s ship) waypoint(v vec) {
	s.v.x += v.x
	s.v.y += v.y
}

func (s ship) fromOrigin() int {
	return abs(s.p.x) + abs(s.p.y)
}

func (s ship) rot(a int) {
	beta := float64(a) * math.Pi / 180

	vx := int(math.Round(math.Cos(beta)*float64(s.v.x) + math.Sin(beta)*float64(s.v.y)))
	vy := int(math.Round(-math.Sin(beta)*float64(s.v.x) + math.Cos(beta)*float64(s.v.y)))

	s.v.x = vx
	s.v.y = vy
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

	s := ship{
		p: &point{0, 0},
		v: &vec{1, 10},
	}

	for _, rawLine := range strings.Split(string(b), "\n") {
		line := strings.TrimSpace(rawLine)

		if line == "" {
			continue
		}

		dir := line[0]
		n, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		switch dir {
		case 'F':
			s.forward(n)
		case 'N':
			s.waypoint(vec{n, 0})
		case 'S':
			s.waypoint(vec{-n, 0})
		case 'W':
			s.waypoint(vec{0, -n})
		case 'E':
			s.waypoint(vec{0, n})
		case 'L':
			s.rot(n)
		case 'R':
			s.rot(-n)
		}
	}

	fmt.Println(s.fromOrigin())
}
