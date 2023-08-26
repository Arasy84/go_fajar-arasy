package main

import (
	"fmt"
)

func munculSekali(angka string) []int {
	counts := make(map[int]int)
	for _, digit := range angka {
		counts[int(digit-'0')]++
	}

	result := []int{}
	for _, digit := range angka {
		if counts[int(digit-'0')] == 1 {
			result = append(result, int(digit-'0'))
		}
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
