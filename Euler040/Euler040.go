package main

import "strconv"

func main() {
	// This is cringe, maybe I'll come back to it later with a cleaner solution
	s := ""
	for i := 0; len(s) <= 1000000; i++ {
		s += strconv.Itoa(i)
	}

	product := numAt(s, 1)
	product *= numAt(s, 10)
	product *= numAt(s, 100)
	product *= numAt(s, 1000)
	product *= numAt(s, 10000)
	product *= numAt(s, 100000)
	product *= numAt(s, 1000000)
	println(product)
}

func numAt(str string, index int) int {
	x, _ := strconv.Atoi(string(str[index]))
	return x
}
