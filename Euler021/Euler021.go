package main

import "math"

func main() {
	sum := 0
	for a := 1; a < 10000; a++ {
		b := d(a)
		if a != b && d(b) == a {
			// Checking for duplicates? too lazy
			sum += a
		}
	}
	println(sum)
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
