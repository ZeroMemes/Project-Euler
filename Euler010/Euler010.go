package main

import "math"

func main() {
	sum := 0

	for i := 1; i < 2000000; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	println(sum)
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
