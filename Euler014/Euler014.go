package main

func main() {
	longest := 0
	starting := 0

	for i := 1; i < 1000000; i++ {
		length := 0
		for j := i; j > 1; length++ {
			j = collatz(j)
		}
		if length > longest {
			longest = length
			starting = i
		}
	}

	println(starting)
}

func collatz(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}
