package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	result := ""
	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			shiftedChar := 'a' + ((char - 'a' + rune(offset)) % 26)
			result += string(shiftedChar)
		} else if char >= 'A' && char <= 'Z' {
			shiftedChar := 'A' + ((char - 'A' + rune(offset)) % 26)
			result += string(shiftedChar)
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}
