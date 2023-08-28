package main

import (
	"fmt"
	"strings"
)

func Compare(A, B string) string {
	if strings.Contains(A, B) {
		return B
	} else if strings.Contains(B, A) {
		return A
	} else {
		return ""
	}
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
