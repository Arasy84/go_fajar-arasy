package main

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}

func factorial(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock()
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	mutex.Unlock()

	fmt.Printf("Factorial of %d is %d\n", n, result)
}

func main() {
	var wg sync.WaitGroup

	numbers := []int{8, 7, 6, 2}

	for _, num := range numbers {
		wg.Add(1)
		go factorial(num, &wg)
	}

	wg.Wait()
}
