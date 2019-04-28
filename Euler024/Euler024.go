package main

import "sort"

func main() {
	permutations := getPermutations("0123456789")
	sort.Strings(permutations)
	println(permutations[999999])
}

func getPermutations(str string) []string {
	permutations := make([]string, 0)
	permutate(&permutations, str, 0)
	return permutations
}

func permutate(permutations *[]string, str string, index int) {
	if index == len(str)-1 {
		*permutations = append(*permutations, str)
		return
	}

	for i := index; i < len(str); i++ {
		// Swap characters
		chars := []rune(str)
		chars[i], chars[index] = chars[index], chars[i]
		permutate(permutations, string(chars), index+1)
	}
}
