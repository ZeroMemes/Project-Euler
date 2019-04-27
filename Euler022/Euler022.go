package main

import (
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("Euler022/names.txt")
	split := strings.Split(string(bytes), ",")
	sort.Strings(split)

	total := 0

	for i := 0; i < len(split); i++ {
		str := strings.Trim(split[i], `"`)
		sum := 0
		for j := 0; j < len(str); j++ {
			sum += int(str[j]) - 'A' + 1
		}
		total += sum * (i + 1)
	}

	println(total)
}
