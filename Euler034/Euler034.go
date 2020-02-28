package main

func main() {
	total := 0

	factorial := make([]int, 10)
	factorial[0] = 1
	for i := 1; i < len(factorial); i++ {
		factorial[i] = factorial[i-1] * i
	}

	for i := 3; i < 1000000; i++ {
		sum := 0

		v := i
		for v > 0 {
			sum += factorial[v%10]
			v /= 10
		}

		if sum == i {
			total += i
		}
	}

	println(total)
}
