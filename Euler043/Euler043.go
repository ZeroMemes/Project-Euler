package main

import "strconv"

func main() {
	sum := 0
	primes := []int{2, 3, 5, 7, 11, 13, 17}
	perms := getPermutations("0123456789")

outer:
	for _, perm := range perms {
		i, _ := strconv.Atoi(perm)
		for j := 0; j < 7; j++ {
			if substringNum(perm, j+1)%primes[j] != 0 {
				continue outer
			}
		}
		sum += i
	}

	println(sum)
}

func substringNum(str string, start int) int {
	i, _ := strconv.Atoi(str[start : start+3])
	return i
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
