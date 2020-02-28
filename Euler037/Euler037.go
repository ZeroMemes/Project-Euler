package main

import "strconv"

func main() {
	sum := 0

outer:
	for i := 10; i < 1000000; i++ {
		if !isPrime(int64(i)) {
			continue
		}
		s := strconv.Itoa(i)
		for j := 1; j < len(s); j++ {
			left, _ := strconv.Atoi(s[j:])
			right, _ := strconv.Atoi(s[:len(s)-j])
			if !isPrime(int64(left)) || !isPrime(int64(right)) {
				continue outer
			}
		}
		sum += i
	}

	println(sum)
}

func isPrime(n int64) bool {
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
