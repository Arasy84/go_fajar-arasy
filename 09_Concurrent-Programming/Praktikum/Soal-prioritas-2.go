package main

import (
	"fmt"
	"unicode"
)

func count_letters(text string) map[rune]int {
	frequencies := make(map[rune]int)

	for _, c := range text {
		if !unicode.IsLetter(c) {
			continue
		}

		frequencies[c]++
	}

	return frequencies
}

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"

	frequencies := count_letters(text)

	for letter, count := range frequencies {
		fmt.Printf("%c: %d\n", letter, count)
	}
}
