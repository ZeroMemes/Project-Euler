package main

func main() {
	size := 20
	println(binomialCoefficient(2*size, size))
}

// method from wikipedia lol
func binomialCoefficient(n, k int) int64 {
	if k < 0 || k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}
	k = min(k, n-k)
	c := int64(1)
	for i := 0; i < k; i++ {
		c = c * (int64(n) - int64(i)) / (int64(i) + 1)
	}
	return c
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
