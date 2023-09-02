package main

import (
	"fmt"
)

func primeX(n int) int {
	if n <= 0 {
		return 0
	}

	count := 0
	num := 1

	for {
		num++
		if isPrime(num) {
			count++
			if count == n {
				return num
			}
		}
	}
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}

	for i := 5; i*i <= num; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}
