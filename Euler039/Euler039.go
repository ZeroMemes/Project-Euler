package main

func main() {
	maxP := 0
	maxSolutions := 0
	for p := 0; p <= 1000; p++ {
		count := 0
		for a := 1; a <= p; a++ {
			for b := 1; b <= p-a; b++ {
				c := p - a - b
				if a*a+b*b == c*c {
					count++
				}
			}
		}
		if count > maxSolutions {
			maxSolutions = count
			maxP = p
		}
	}
	println(maxP)
}
