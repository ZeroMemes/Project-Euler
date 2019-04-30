package main

import "math/big"

func main() {
	calcs := make(map[string]struct{})

	for a := 2; a <= 100; a++ {
		bigA := big.NewInt(int64(a))
		for b := 2; b <= 100; b++ {
			bigB := big.NewInt(int64(b))
			calcs[big.NewInt(0).Exp(bigA, bigB, nil).String()] = struct{}{}
		}
	}

	println(len(calcs))
}
