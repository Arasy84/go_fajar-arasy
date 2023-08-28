package main

import "fmt"

func findMinMax(numbers []int) (int, int) {
	if len(numbers) == 0 {
		return 0, 0
	}

	max := numbers[0]
	min := numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return max, min
}

func main() {
	var inputNumbers [6]int

	fmt.Println("Input:")
	for i := 0; i < 6; i++ {
		fmt.Scan(&inputNumbers[i])
	}

	max, min := findMinMax(inputNumbers[:])

	fmt.Printf("%d is maximum number\n", max)
	fmt.Printf("%d is minimum number\n", min)
}
