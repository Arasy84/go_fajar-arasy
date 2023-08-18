package main

import "fmt"

func main() {
	var tinggi, sisiAtas, sisiBawah float64

	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scanln(&tinggi)

	fmt.Print("Masukkan panjang sisi atas trapesium: ")
	fmt.Scanln(&sisiAtas)

	fmt.Print("Masukkan panjang sisi bawah trapesium: ")
	fmt.Scanln(&sisiBawah)

	luas := (sisiAtas + sisiBawah) * tinggi / 2

	fmt.Println("Luas trapesium:", luas)
}
