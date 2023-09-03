package main

import "fmt"

func fibonacci(number int) int {
	var memo = make(map[int]int)
	if number <= 1 {
		return number
	}

	if val, found := memo[number]; found {
		return val
	}

	result := fibonacci(number-1) + fibonacci(number-2)
	memo[number] = result
	return result
}

func main() {
	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(9))
	fmt.Println(fibonacci(10))
	fmt.Println(fibonacci(12))
}
