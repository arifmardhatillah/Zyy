package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(calculateCircle(7))

}

func calculateCircle(diameter float64) (float64, float64) {
	luas := math.Pi * math.Pow(diameter/2, 2)
	keliling := math.Pi * diameter
	// return 2 value
	return keliling, luas
}
