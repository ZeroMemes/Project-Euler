package main

import (
	"math"
)

func main() {
	all := 0

	// Cache this because it zooms. Also makes the next portion look a lot cleaner
	powers := [10]int{}
	for i := 0; i < 10; i++ {
		powers[i] = int(math.Pow(float64(i), 5))
	}

	// 5*9^5 should be a sensible upper bound, idk it's kind of arbitrary
	for i := 2; i < 5*int(math.Pow(9, 5)); i++ {
		sum := 0
		n := i
		for n > 0 {
			sum += powers[n%10]
			n /= 10
		}
		if sum == i {
			all += i
		}
	}

	println(all)
}
