package main

func main() {
	sum := 0
	for a := 1; a < 10000; a++ {
		b := d(a)
		if a != b && d(b) == a {
			// Checking for duplicates? too lazy
			sum += a
		}
	}
	println(sum)
}

func d(n int) int {
	sum := 0
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}
