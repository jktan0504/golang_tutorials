package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

// struct definition
type User struct {
	firstName string
	lastName string
	birthDate string
	createdDate time.Time
} 

func (user *User) outputDetails() {
	fmt.Printf("MY name is %v %v {born on %v}", user.firstName, user.lastName, user.birthDate)
}


func outputUserDetails(user *User) {
	fmt.Printf("MY name is %v %v {born on %v}", user.firstName, user.lastName, user.birthDate)
}

// * Type = pointer, e.g. *User
func NewUser(fName string, lName string, bDate string) *User {
	return &User{
		firstName: fName,
		lastName: lName,
		birthDate: bDate,
		createdDate: time.Now(),
	}
}

var reader = bufio.NewReader(os.Stdin)

func main() {

	var newUser User

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// adding * to pointer to get value
	newUser = *NewUser(firstName, lastName, birthdate)
	outputUserDetails(&newUser)
	newUser.outputDetails()
}



func getUserData(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')

	var cleanedInput string;

	if runtime.GOOS == "windows" {
		cleanedInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		cleanedInput = strings.Replace(userInput, "\n", "", -1)
	}
	return cleanedInput

}