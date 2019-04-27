package main

import "math"

func main() {
	const bound = 28123

	abundantNumbers := make([]bool, bound)
	for i := 1; i <= bound; i++ {
		abundantNumbers[i-1] = isAbundant(i)
	}

	sum := 0

outer:
	for i := 1; i < bound; i++ {
		for j := 1; j < i; j++ {
			if abundantNumbers[j-1] && abundantNumbers[(i-j)-1] {
				continue outer
			}
		}
		sum += i
	}

	println(sum)
}

func isAbundant(n int) bool {
	return d(n) > n
}

func d(n int) int {
	sum := 1
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if i == n/i {
				sum += i
			} else {
				sum += i + n/i
			}
		}
	}
	return sum
}
