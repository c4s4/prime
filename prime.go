package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func divide(i int) (int, int) {
	r := int(math.Sqrt(float64(i)))
	if i%2 == 0 {
		return 2, i / 2
	}
	for d := 3; d <= r; d++ {
		if i%d == 0 {
			return d, i / d
		}
	}
	return i, 0
}

func split(i int) []int {
	var factors []int
	for {
		j, k := divide(i)
		factors = append(factors, j)
		if k == 0 {
			return factors
		}
		i = k
	}
}

func format(i int, f []int) string {
	var g []string
	for _, h := range f {
		g = append(g, strconv.Itoa(h))
	}
	return fmt.Sprintf("%d = %s", i, strings.Join(g, " x "))
}

func main() {
	for _, n := range os.Args[1:] {
		i, err := strconv.Atoi(n)
		if err == nil {
			fmt.Println(format(i, split(i)))
		} else {
			fmt.Printf("%s n'est pas un entier\n", n)
		}
	}
}
