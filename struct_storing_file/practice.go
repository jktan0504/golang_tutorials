package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Product
type Product struct {
	id string
	title string
	description string
	price float64
}

func NewProduct(id string, title string, description string, price float64) (product *Product) {
	return &Product {
		id,
		title,
		description,
		price,
	}
}

func (product *Product) outputDetails() {
	fmt.Println(product)
	fmt.Printf("This book is %v which price %.2f", product.title, product.price)
}

func (prod *Product) store() {
	file, _ := os.Create(prod.id + ".txt")
	content := fmt.Sprintf(
		`ID: %v\n
			Title: %v\nDescription: %v\n
			Price: %.2f\n\n
		`, prod.id, prod.title, prod.description, prod.price)
	file.WriteString(content)
	file.Close()
}

func getProduct() *Product {
	fmt.Println("Please enter the product data")
	fmt.Println("==============================")

	reader := bufio.NewReader(os.Stdin)
	idInput := readUserInput(reader, "ID: ")
	titleInput := readUserInput(reader, "Title: ")
	desInput := readUserInput(reader, "Description: ")
	priceInput := readUserInput(reader, "Price: ")
	priceValue, _ := strconv.ParseFloat(priceInput, 64)
	
	product := NewProduct(idInput, titleInput, desInput, priceValue)
	return product
}

func readUserInput(reader *bufio.Reader, promtText string) (string) {
	fmt.Println(promtText)
	valueInput, _  := reader.ReadString('\n')
	valueInput = strings.Replace(valueInput, "\n", "", -1)
	return valueInput
}

func main() {
	createdProduct := getProduct()
	createdProduct.outputDetails()
	createdProduct.store()
}

// Your Tasks
// 1) Create a new type / data structure for storing product data (e.g. about a book)
//		The data structure should contain these fields:
//		- ID
//		- Title / Name
//		- Short description
//		- price (number without currency, we'll just assume USD)
// 2) Create concrete instances of this data type in two different ways:
//		- Directly inside of the main function
//		- Inside of main, by using a "creation helper function"
//		(use any values for title etc. of your choice)
//		Output (print) the created data structures in the command line (in the main function)
// 3) Add a "connected function" that outputs the data + call that function from inside "main"
// 4) Change the program to fetch user input values for the different data fields
//		and create only one concrete instance with that entered data.
//		Output that instance data (via the connected function) then.
// 5) Bonus: Add a connected "store" function that writes that data into a file
//		The filename should be the unique id, the function should be called at the
//		end of main.