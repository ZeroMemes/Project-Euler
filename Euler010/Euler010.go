package main

func main() {
	sum := 0

	for i := 1; i < 2000000; i++ {
		if isPrime(int64(i)) {
			sum += i
		}
	}

	println(sum)
}

func isPrime(n int64) bool {
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
