package main

func main() {
	d, longest := 0, 0

	for i := 1; i < 1000; i++ {
		c := cycle(i)
		if c > longest {
			longest = c
			d = i
		}
	}

	println(d)
}

func cycle(n int) int {
	digits := 0
	div := 1
	remainders := make(map[int]bool)
	for ; !remainders[div] && div != 0; digits++ {
		remainders[div] = true
		div *= 10
		div %= n
	}
	return digits
}
