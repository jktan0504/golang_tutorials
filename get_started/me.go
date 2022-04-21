package main

import (
	"fmt"
)

func me() {
	firstName := "JK"
	var lastname string = "Tan"
	birthYear := 1992
	var currentYear int = 2022
	age := currentYear - birthYear

	fmt.Println(firstName + " " + lastname)
	fmt.Println(age)

}
