package main

import (
	"fmt"
	"math"
)

func main() {
	largest := 0

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			p := i * j
			if isPalindrome(p) {
				largest = max(p, largest)
			}
		}
	}

	println(largest)
}

func isPalindrome(n int) bool {
	str := fmt.Sprintf("%d", n)
	for i := 0; i < floor(len(str)/2); i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func floor(a int) int {
	return int(math.Floor(float64(a)))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
