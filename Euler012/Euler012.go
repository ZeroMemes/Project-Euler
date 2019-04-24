package main

func main() {
	for i := 1; ; i++ {
		triangle := triangle(i)
		if divisors(triangle) > 500 {
			println(triangle)
			break
		}
	}
}

func triangle(n int) int {
	if n == 1 {
		return 1
	}
	return n + triangle(n-1)
}

func divisors(n int) int {
	if n == 1 {
		return 1
	}

	factors := 0
	for i := 1; i*i < n; i++ {
		if n%i == 0 {
			factors += 2
		}
	}
	return factors
}
