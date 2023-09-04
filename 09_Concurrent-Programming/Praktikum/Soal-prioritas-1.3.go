package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10)

	go generateMultiplesOfThree(ch)

	for i := 0; i < 10; i++ {
		multiple := <-ch
		fmt.Printf("Kelipatan 3 ke-%d: %d\n", i+1, multiple)
	}
}

func generateMultiplesOfThree(ch chan int) {
	for i := 1; i <= 30; i++ {
		ch <- i * 3
	}
	close(ch)
}
