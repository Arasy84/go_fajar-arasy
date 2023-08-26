package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PairSum(arr []int, target int) []int {
	left := 0
	right := len(arr) - 1

	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan jumlah elemen dalam array: ")
	numElementsInput, _ := reader.ReadString('\n')
	numElementsInput = strings.TrimSpace(numElementsInput)
	numElements, _ := strconv.Atoi(numElementsInput)

	arr := make([][]int, numElements)

	for i := 0; i < numElements; i++ {
		fmt.Printf("Masukkan elemen-elemen array ke-%d (pisahkan dengan spasi): ", i+1)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		elements := strings.Split(input, " ")

		arr[i] = make([]int, len(elements))
		for j, element := range elements {
			num, _ := strconv.Atoi(element)
			arr[i][j] = num
		}
	}

	for i, elementArray := range arr {
		fmt.Printf("Array ke-%d: %v\n", i+1, elementArray)

		fmt.Print("Masukkan target sum untuk array ini: ")
		targetInput, _ := reader.ReadString('\n')
		targetInput = strings.TrimSpace(targetInput)
		target, _ := strconv.Atoi(targetInput)

		result := PairSum(elementArray, target)

		if len(result) == 0 {
			fmt.Println("Tidak ditemukan pasangan yang sesuai dalam array ini.")
		} else {
			fmt.Printf("Indeks pasangan dalam array ini: %v\n", result)
		}
	}
}
