package main

import "fmt"

func main() {
	// pointers are variable to store the address of a value instead of value itself
	age := 30

	// pointer
	// var userAge *int
	// userAge = &age
	// Pointer at an integer value
	userAge := &age

	// using & to get address
	// using * to get value
	fmt.Printf("age is %v and pointer is %v (%v)\n", age, userAge, *userAge)

	// re-assign value to the pointer's value
	*userAge = 20
	fmt.Printf("re-assigned age is %v and pointer is %v (%v)", age, userAge, *userAge)

	doubleMyAgeValue := doubleMyAge(userAge)
	fmt.Println(doubleMyAgeValue)
}

func doubleMyAge(pointerAge *int) (int){
	return *pointerAge * 2
}
