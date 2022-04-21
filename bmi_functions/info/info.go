package info

import "fmt"

const mainTitle = "\n\nBMI Calculator"
const separator = "------------------"
const WeightPrompt = "Please enter your weight (KG) :"
const HeightPrompt = "Please enter your height (m)  :"

func PrintGreetingMessages() {
	fmt.Println(mainTitle)
	fmt.Println(separator)
}

func PrintResult(weight float64, height float64, bmi float64) {
	fmt.Printf("Weight: %.2f\nHeight: %.2f\nBMI: %.2f", weight, height, bmi)
}
