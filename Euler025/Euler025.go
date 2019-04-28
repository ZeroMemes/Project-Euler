package main

import "math/big"

func main() {
	n1 := big.NewInt(1)
	n2 := big.NewInt(1)
	index := 2
	for len(n2.String()) < 1000 {
		n1, n2 = n2, n1.Add(n1, n2)
		index++
	}
	println(index)
}
