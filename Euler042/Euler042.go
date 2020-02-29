package main

import (
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("Euler042/words.txt")
	split := strings.Split(string(bytes), ",")

	total := 0

	for _, s := range split {
		s = strings.Trim(s, `"`)
		n := 0
		for _, c := range s {
			n += int(c - 64)
		}

		// If t^-1(n) is an integer, then it's a valid triangle number
		ti := 0.5 * (-1 + math.Sqrt(float64(8*n+1)))
		if ti == float64(int64(ti)) {
			total++
		}
	}

	println(total)
}
