package main

import (
	"fmt"
	"math/rand"
)

func main() {
	number1, number2 := randomNumberGenerator()
	printNumber(number1, number2, add(number1, number2))
}

// Random Number
func randomNumberGenerator() (int, int) {
	num1 := rand.Intn(10)
	num2 := rand.Intn(10)
	return num1, num2
}

// Sum two numbers
// func add(num1 int, num2 int) int {
// 	sum := num1 + num2
// 	return sum
// }
// Using named returned values
func add(num1 int, num2 int) (sum int) {
	sum = num1 + num2
	return
}

// Print Result
func printNumber(number1 int, number2 int, result int) {
	fmt.Printf("%v + %v = %v", number1, number2, result)
}
