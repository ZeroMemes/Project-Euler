package main

import "strconv"

func main() {
	sum := 0

	for i := 1; i < 1000000; i++ {
		base10 := strconv.Itoa(i)
		base2 := strconv.FormatInt(int64(i), 2)
		if isPalindromic(base10) && isPalindromic(base2) {
			sum += i
		}
	}

	println(sum)
}

func isPalindromic(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
