package main

import (
	"strconv"
)

func main() {
	maximum := 0
	s := ""

	for i := 1; i <= 9; i++ {
		s += strconv.Itoa(i)
		perms := getPermutations(s)
		for _, perm := range perms {
			i, _ := strconv.Atoi(perm)
			if isPrime(int64(i)) {
				maximum = max(maximum, i)
			}
		}
	}

	println(maximum)
}

func getPermutations(str string) []string {
	var perms []string
	permutations(str, 0, &perms)
	return perms
}

func permutations(str string, index int, perms *[]string) {
	if index == len(str)-1 {
		*perms = append(*perms, str)
	} else {
		for i := index; i < len(str); i++ {
			permutations(swap(str, index, i), index+1, perms)
		}
	}
}

func swap(str string, i1, i2 int) string {
	ret := []rune(str)
	ret[i1], ret[i2] = ret[i2], ret[i1]
	return string(ret)
}

func isPrime(n int64) bool {
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
