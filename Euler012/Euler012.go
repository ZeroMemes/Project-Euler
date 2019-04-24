package main

func main() {
	num := 1
	for i := 1; ; {
		if divisors(num) > 500 {
			println(num)
			break
		}
		i++
		num += i
	}
}

func divisors(n int) int {
	if n == 1 {
		return 1
	}

	factors := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			factors += 2
		}
	}
	return factors
}
