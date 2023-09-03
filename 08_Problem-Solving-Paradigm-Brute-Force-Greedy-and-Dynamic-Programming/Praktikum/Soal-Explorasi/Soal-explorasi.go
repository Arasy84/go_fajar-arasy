package main

import "fmt"

func convertToRoman(num int) string {
	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"}, {100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"}, {10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	romanNumeral := ""

	for num > 0 {
		for _, numeral := range romanNumerals {
			if num >= numeral.Value {
				romanNumeral += numeral.Symbol
				num -= numeral.Value
				break
			}
		}
	}

	return romanNumeral
}

func main() {
	inputNums := []int{4, 9, 23, 2021, 1646}

	for _, inputNum := range inputNums {
		outputRoman := convertToRoman(inputNum)
		fmt.Printf("Input: %d\n", inputNum)
		fmt.Printf("Output: %s\n", outputRoman)
		fmt.Println()
	}
}
