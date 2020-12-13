package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const file = "data.txt"

func main() {
	fileName := file

	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	n := make([]*big.Int, 0)
	a := make([]*big.Int, 0)

	l := int64(0)

	for i, rawLine := range strings.Split(string(b), "\n") {
		line := strings.TrimSpace(rawLine)

		if line == "" {
			continue
		}

		if i == 1 {
			for j, b := range strings.Split(line, ",") {
				l++
				if b != "x" {
					m, err := strconv.ParseInt(b, 0, 64)
					if err != nil {
						panic(err)
					}

					a = append(a, big.NewInt(m))
					n = append(n, big.NewInt(m-int64(j)))
				}
			}
		}
	}

	fmt.Println(a)
	fmt.Println(n)

	if sol, err := crt(n, a); err != nil {
		panic(err)
	} else {
		t := sol
		for i := int64(0); i <= l+1; i++ {
			fmt.Printf("%d\t", t.Int64())
			for j := range a {
				tmp := new(big.Int).Mod(t.Add(t, big.NewInt(i)), a[j])
				fmt.Printf("%3d ", tmp.Int64())
			}
			fmt.Printf(" < %d\n", i)
		}

		fmt.Println("sol:", sol)
	}
}

var one = big.NewInt(1)

func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}
