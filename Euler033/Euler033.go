package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	num := 1
	den := 1

	for x := 10; x < 100; x++ {
		for y := x + 1; y < 100; y++ {
			xs := strconv.Itoa(x)
			ys := strconv.Itoa(y)

			// Exclude fractions that work by reducing the magnitude
			if x%10 == 0 && y%10 == 0 {
				continue
			}

			// Check that at least one of the numbers matches
			if !strings.Contains(xs, string(ys[0])) && !strings.Contains(xs, string(ys[1])) {
				continue
			}

			// Remove all duplicate numbers
			xr := regexp.MustCompile("[" + xs + "]")
			yr := regexp.MustCompile("[" + ys + "]")
			xs = yr.ReplaceAllString(xs, "")
			ys = xr.ReplaceAllString(ys, "")

			// Ensure that the resulting numerator and denominator are valid to consider
			if len(xs) == 0 || len(ys) == 0 || xs == "0" || ys == "0" {
				continue
			}

			// Check that the resulting fraction retains the original value
			xi, _ := strconv.Atoi(xs)
			yi, _ := strconv.Atoi(ys)
			if float64(x)/float64(y) == float64(xi)/float64(yi) {
				num *= xi
				den *= yi
			}
		}
	}

	fmt.Printf("%d/%d", num, den)
}
