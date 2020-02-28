package main

import "strconv"

func main() {
	count := 0
outer:
	for i := 2; i < 1000000; i++ {
		if !isPrime(int64(i)) {
			continue
		}
		s := strconv.Itoa(i)
		for j := 1; j < len(s); j++ {
			v, _ := strconv.Atoi(s[j:] + s[0:j])
			if !isPrime(int64(v)) {
				continue outer
			}
		}
		count++
	}
	println(count)
}

func isPrime(n int64) bool {
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
