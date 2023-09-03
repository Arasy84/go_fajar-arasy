package main

import (
	"fmt"
	"strconv"
)

func binaryRepresentation(n int) []string {
	ans := make([]string, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = strconv.FormatInt(int64(i), 2)
	}
	return ans
}

func main() {
	n1 := 2
	result := binaryRepresentation(n1)
	fmt.Printf("Input: n = %d\n", n1)
	fmt.Println(result)

	n2 := 3
	result = binaryRepresentation(n2)
	fmt.Printf("\nInput: n = %d\n", n2)
	fmt.Println(result)
}
