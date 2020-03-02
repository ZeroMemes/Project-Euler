package main

import "math"

func main() {
	minimum := math.MaxInt32
	for k := 2; k <= 10000; k++ {
		for j := 1; j < k; j++ {
			pk := P(k)
			pj := P(j)
			if isPentagonal(pk+pj) && isPentagonal(pk-pj) {
				minimum = min(minimum, pk-pj)
			}
		}
	}
	println(minimum)
}

func isPentagonal(n int) bool {
	// inverse of P(n)
	p := (1.0 / 6.0) * (1 + math.Sqrt(float64(24*n+1)))
	return p == float64(int64(p))
}

func P(n int) int {
	return n * (3*n - 1) / 2
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
