package main

import "fmt"

func power(x int, n int) int {
	if n == 0 {
		return 1
	} else if n%2 == 0 {
		return power(x*x, n/2)
	} else {
		return x * power(x, n-1)
	}
}

func main() {
	fmt.Println(power(2, 3))
	fmt.Println(power(7, 2))
}
