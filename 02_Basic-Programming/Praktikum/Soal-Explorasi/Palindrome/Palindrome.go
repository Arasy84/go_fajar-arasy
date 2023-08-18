package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(str string) bool {
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, " ", "") // Menghapus spasi
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Apakah palindrome?")
	fmt.Print("Masukkan kata: ")

	var input string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	if isPalindrome(input) {
		fmt.Println("captured:", input)
		fmt.Println("Palindrome")
	} else {
		fmt.Println("captured:", input)
		fmt.Println("Bukan palindrome")
	}
}
