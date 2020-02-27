package main

import (
	"strconv"
)

func main() {
	solutions := make(map[int]struct{})

	var perms []string
	permutations("123456789", 0, &perms)

	for _, perm := range perms {
		c, _ := strconv.Atoi(perm[5:9])

		a, _ := strconv.Atoi(perm[0:1])
		b, _ := strconv.Atoi(perm[1:5])
		if a*b == c {
			solutions[c] = struct{}{}
		}

		a, _ = strconv.Atoi(perm[0:2])
		b, _ = strconv.Atoi(perm[2:5])
		if a*b == c {
			solutions[c] = struct{}{}
		}
	}

	sum := 0
	for solution := range solutions {
		sum += solution
	}
	println(sum)
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

	c := ret[i1]
	ret[i1] = ret[i2]
	ret[i2] = c

	return string(ret)
}
