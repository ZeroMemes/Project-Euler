package main

import (
	"math/big"
	"strconv"
)

func main() {
	base := big.NewInt(2)
	base.Exp(base, big.NewInt(1000), nil)
	str := base.String()

	sum := 0
	for i := 0; i < len(str); i++ {
		num, _ := strconv.Atoi(string(str[i]))
		sum += num
	}

	println(sum)
}
