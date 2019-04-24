package main

import "math"

func main() {
	sum := 0

	f := float64(1)
	for f < 4000000 {
		f = math.Round(f * math.Phi)
		if int(f)%2 == 0 {
			sum += int(f)
		}
	}

	println(sum)
}
