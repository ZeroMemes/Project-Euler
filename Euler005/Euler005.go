package main

func main() {
	num := 0

outer:
	for true {
		num++
		for i := 1; i <= 20; i++ {
			if num%i != 0 {
				continue outer
			}
		}
		break
	}

	println(num)
}
