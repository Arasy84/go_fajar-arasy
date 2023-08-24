package main

import (
	"fmt"
	"math"
	"time"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	numbers := []int{1000000007, 1500450271, 13, 17, 20, 35}

	for _, num := range numbers {
		start := time.Now()
		isPrime := isPrime(num)
		elapsed := time.Since(start)

		fmt.Printf("Hasil %v Waktu (%v seconds)\n", isPrime, elapsed.Seconds())
	}
}
