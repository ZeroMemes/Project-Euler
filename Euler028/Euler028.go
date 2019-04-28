package main

func main() {
	sum := 1
	size := 1001

	for i := 0; i < size/2; i++ {
		size := (i + 1) * 2
		base := size + 1
		square := base * base
		sum += square*4 - size*6
	}

	println(sum)
}
