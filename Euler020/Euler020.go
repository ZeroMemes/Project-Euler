package main

import (
	"math/big"
	"strconv"
)

func main() {
	result := factorial(100)
	str := result.String()

	sum := 0
	for i := 0; i < len(str); i++ {
		num, _ := strconv.Atoi(string(str[i]))
		sum += num
	}

	println(sum)
}

func factorial(n int) *big.Int {
	result := big.NewInt(int64(n))
	for i := 1; i < n-1; i++ {
		result = result.Mul(result, big.NewInt(int64(n-i)))
	}
	return result
}
