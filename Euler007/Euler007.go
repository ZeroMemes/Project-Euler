package main

import "math"

func main() {
	primes := 1
	num := 1

	for primes < 10001 {
		num += 2
		if isPrime(num) {
			primes++
		}
	}

	println(num)
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
