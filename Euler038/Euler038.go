package main

import (
	"math"
	"strconv"
)

func main() {
	max := 0

	for i := 1; i < 1000000; i++ {
		s := ""
		for j := 1; j <= 9; j++ {
			s += strconv.Itoa(i * j)
			v, _ := strconv.Atoi(s)

			if len(s) == 9 && isPandigital1to9(v) {
				max = int(math.Max(float64(max), float64(v)))
			}
		}
	}

	println(max)
}

func isPandigital1to9(i int) bool {
	digits := make([]int, 9)
	v := i
	for v > 0 {
		d := v % 10
		if d == 0 {
			return false
		}
		digits[d-1]++
		v /= 10
	}
	for _, d := range digits {
		if d != 1 {
			return false
		}
	}
	return true
}
