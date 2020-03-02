package main

import (
	"math"
)

type gamerMap map[int]bool

var primes = make(gamerMap)
var squares = make(gamerMap)

func main() {
outer:
	for i := 9; true; i += 2 {
		if isPrime(i) {
			continue
		}

		for j := i - 2; j > 0; j -= 2 {
			if isPrime(j) && isSquare((i-j)/2) {
				continue outer
			} else if j <= 1 {
				println(i)
				break outer
			}
		}
	}
}

func (gm *gamerMap) computeIfAbsent(key int, eval func(int) bool) bool {
	if _, ok := (*gm)[key]; !ok {
		(*gm)[key] = eval(key)
	}
	return (*gm)[key]
}

func isPrime(n int) bool {
	return primes.computeIfAbsent(n, func(i int) bool {
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return n > 1
	})
}

func isSquare(n int) bool {
	return squares.computeIfAbsent(n, func(i int) bool {
		root := math.Sqrt(float64(i))
		return float64(int64(root)) == root
	})
}
