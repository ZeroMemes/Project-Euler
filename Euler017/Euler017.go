package main

import (
	"regexp"
	"strings"
)

var onesNames = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

var teensNames = map[int]string{
	1: "eleven",
	2: "twelve",
	3: "thirteen",
	4: "fourteen",
	5: "fifteen",
	6: "sixteen",
	7: "seventeen",
	8: "eighteen",
	9: "nineteen",
}

var tensNames = map[int]string{
	1: "ten",
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

func main() {
	regex := regexp.MustCompile("[^a-z]+")
	length := 0
	for i := 1; i <= 1000; i++ {
		length += len(regex.ReplaceAllString(toString(i), ""))
	}
	println(length)
}

func toString(n int) string {
	thousands := n / 1000 % 10
	hundreds := n / 100 % 10
	tens := n / 10 % 10
	ones := n % 10

	var elements []string

	if thousands > 0 {
		elements = append(elements, onesNames[thousands]+" thousand")
	}

	if hundreds > 0 {
		if thousands > 0 && (tens > 0 || ones > 0) {
			elements[0] += ","
		}
		elements = append(elements, onesNames[hundreds]+" hundred")
	}

	if (thousands > 0 || hundreds > 0) && (tens > 0 || ones > 0) {
		elements = append(elements, "and")
	}

	if tens >= 2 || tens == 1 && ones == 0 {
		str := ""
		if ones > 0 {
			str += "-" + onesNames[ones]
		}
		elements = append(elements, tensNames[tens]+str)
	} else if tens == 1 {
		elements = append(elements, teensNames[ones])
	} else if ones > 0 {
		elements = append(elements, onesNames[ones])
	}

	return strings.Join(elements, " ")
}
