package main

func main() {
	product := 0
	greatest := 0

	for b := 0; b <= 1000; b++ {
		// We know that b must be prime, because we're starting with
		// n = 0, and 2 terms of the quadratic function are in terms of n
		if !isPrime(int64(b)) {
			continue
		}
		for a := -999; a < 1000; a++ {
			formula := getFormula(a, b)
			n := 0
			for isPrime(int64(formula(n))) {
				n++
			}
			if n > greatest {
				product = a * b
				greatest = n
			}
		}
	}

	println(product)
}

func getFormula(a, b int) func(int) int {
	return func(n int) int {
		return n*n + a*n + b
	}
}

func isPrime(n int64) bool {
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
