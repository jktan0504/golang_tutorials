package math

import "fmt"

func math() {
	// pi
	var pi float64 = 3.142
	radius := 6
	circumference := pi * float64(radius)

	fmt.Println(circumference)
	// %.2f 2 decimals points only
	fmt.Printf("The radius is %v, and its circumference is %.2f, ", radius, circumference)
}
