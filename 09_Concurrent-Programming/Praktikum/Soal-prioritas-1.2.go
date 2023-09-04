package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go generateMultiplesOfThree(ch, &wg)

	wg.Add(1)
	go printMultiples(ch, &wg)

	wg.Wait()
}

func generateMultiplesOfThree(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		ch <- i * 3
	}
	close(ch)
}

func printMultiples(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for multiple := range ch {
		fmt.Printf("Kelipatan 3: %d\n", multiple)
	}
}
