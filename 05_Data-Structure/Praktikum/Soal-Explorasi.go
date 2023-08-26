package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan ukuran matriks (n): ")
	fmt.Scan(&n)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Printf("Masukkan elemen matriks [%d][%d]: ", i, j)
			fmt.Scan(&matrix[i][j])
		}
	}

	sumKirikeKanan := 0
	for i := 0; i < n; i++ {
		sumKirikeKanan += matrix[i][i]
	}

	sumKanankeKiri := 0
	for i := 0; i < n; i++ {
		sumKanankeKiri += matrix[i][n-1-i]
	}

	Perbedaan := int(math.Abs(float64(sumKirikeKanan - sumKanankeKiri)))

	fmt.Printf("Diagonal kiri ke kanan: %d\n", sumKirikeKanan)
	fmt.Printf("Diagonal kanan ke kiri: %d\n", sumKanankeKiri)
	fmt.Printf("Perbedaan mutlak: %d\n", Perbedaan)
}
