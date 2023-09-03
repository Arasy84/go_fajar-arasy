package main

import "fmt"

func generate(numRows int) [][]int {
	pascalTriangle := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row[j] = 1
			} else {
				row[j] = pascalTriangle[i-1][j-1] + pascalTriangle[i-1][j]
			}
		}
		pascalTriangle[i] = row
	}
	return pascalTriangle
}

func main() {
	numRows := 5
	result := generate(numRows)
	fmt.Println(result)
}
