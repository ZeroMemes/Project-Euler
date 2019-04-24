package main

func main() {
	sumOfSquares := 0
	sum := 0

	for i := 1; i <= 100; i++ {
		sumOfSquares += i * i
		sum += i
	}

	diff := (sum * sum) - sumOfSquares
	println(diff)
}
