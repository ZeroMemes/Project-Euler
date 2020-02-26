package main

var values = []int{200, 100, 50, 20, 10, 5, 2, 1}

func main() {
	total := 0
	foo(&total, 0, 200)
	println(total)
}

func foo(total *int, value, remaining int) {
	for i := 0; i <= remaining/values[value]; i++ {
		nextRemaining := remaining - i*values[value]
		if nextRemaining == 0 || value == len(values)-2 {
			*total++
		} else {
			foo(total, value+1, nextRemaining)
		}
	}
}
