package main

import (
	"fmt" // fmt format package

	"github.com/jktan0504/get_started/variables"
)

func main() {
	// string type
	// variable name type
	// var greetingText string
	// greetingText = "Hello World from Golang !!!"
	// or
	// var greetingText string = "Hello World from Golang !!!"
	// or
	greetingText := variables.GreetingText
	greetingText2 := "JK"

	// int type
	luckyNumber := 199
	anotherNumber := luckyNumber * 6

	fmt.Println(variables.GreetingText)
	fmt.Println(luckyNumber)
	fmt.Println(anotherNumber)

	// float32 & float64
	var decimalNumber float64 = 2.666
	var newDecimalNumber float64 = float64(luckyNumber) / 3
	var smallDecimalNumber float32 = float32(newDecimalNumber)

	fmt.Println(decimalNumber)
	fmt.Println(newDecimalNumber)
	fmt.Println(smallDecimalNumber)

	// rune
	var emoji rune = '$'

	fmt.Println(emoji)         // 36
	fmt.Println(string(emoji)) // $

	// byte
	var firstByte byte = 'a'
	fmt.Println(firstByte)
	fmt.Println(string(firstByte))

	// string operator
	fmt.Println(greetingText + " " + greetingText2)
	// fmt.Println(int("9") + 1)

	fmt.Printf("Greeting: %v , from %v (Type: %T)", greetingText, luckyNumber, luckyNumber)

	abc := fmt.Sprintf("g1: %v", greetingText)
	fmt.Println(abc)

	// Note for String
	// strings.Replace(text, "\n", "", -1) mac or linux
	// strings.Replace(text, "\r\n", "", -1) window
}
