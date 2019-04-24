package main

func main() {
	for c := 1; c < 1000; c++ {
		for b := 1; b < c; b++ {
			for a := 1; a < b; a++ {
				if a+b+c == 1000 && a*a+b*b == c*c {
					println(a * b * c)
					return
				}
			}
		}
	}
}
