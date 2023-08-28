package main

import (
	"fmt"
	"strconv"
)

type Car struct {
	carType string
	fuelin  float64
}

func (c Car) estimateDistance() float64 {
	fuelConsumption := 0.0

	switch c.carType {
	case "suv":
		fuelConsumption = 1.5
	case "sedan":
		fuelConsumption = 1.7
	case "sports":
		fuelConsumption = 1.2
	case "mpv":
		fuelConsumption = 1.8
	default:
		fmt.Println("Invalid car type.")
		return 0.0
	}

	estimatedDistance := c.fuelin * fuelConsumption
	return estimatedDistance
}

func main() {
	var carType string
	var fuelinStr string

	fmt.Print("Masukkan jenis mobil: ")
	fmt.Scan(&carType)

	fmt.Print("Masukkan jumlah bahan bakar (dalam liter): ")
	fmt.Scan(&fuelinStr)

	fuelin, err := strconv.ParseFloat(fuelinStr, 64)
	if err != nil {
		fmt.Println("Input tidak valid. Silakan masukkan angka.")
		return
	}

	myCar := Car{carType: carType, fuelin: fuelin}
	estimatedDistance := myCar.estimateDistance()

	fmt.Printf("Car type: %s, est. distance: %.2f km\n", myCar.carType, estimatedDistance)
}
