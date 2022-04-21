package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jktan0504/bmi/info"
)

func main() {
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)
	// Input
	fmt.Println(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')
	// save input
	weightInput = strings.Replace(weightInput, "\r\n", "", -1)

	fmt.Println(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')
	heightInput = strings.Replace(heightInput, "\r\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	// bmi
	// formula
	bmi := (weight / (height * height)) * 703

	fmt.Printf("\nWeight = %v\nHeight = %v\nBMI: %.2f", weightInput, heightInput, bmi)

}
