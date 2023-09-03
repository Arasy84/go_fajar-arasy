package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) {
	for x := 1; x <= a; x++ {
		for y := 1; y <= a; y++ {
			z := a - x - y
			if x != y && y != z && x != z && x*y*z == b && x*x+y*y+z*z == c {
				fmt.Printf("%d %d %d\n", x, y, z)
				return
			}
		}
	}
	fmt.Println("No Solution")
}

func main() {
	SimpleEquations(1, 2, 3)
	SimpleEquations(6, 6, 14)
}
