package main

import "math"

func main() {
	t := 286
	for true {
		val := T(t)
		if isPentagonal(val) && isHexagonal(val) {
			println(val)
			break
		}
		t++
	}
}

func T(n int) int {
	return n * (n + 1) / 2
}

func isPentagonal(n int) bool {
	// inverse of P(n)
	p := (1.0 / 6.0) * (1 + math.Sqrt(float64(24*n+1)))
	return p == float64(int64(p))
}

func isHexagonal(n int) bool {
	// inverse of H(n)
	h := 0.25 * (1 + math.Sqrt(float64(8*n+1)))
	return h == float64(int64(h))
}
