package main

import "fmt"

// Constants values can not be changed during the time of execution unlike variables
const LoginToken string = "jfkhljlfuid" // Public

func main() {
	// String variables
	var name string = "Daniel"
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T \n", name) // check variale type

	// Booleans data types
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// uint8 data types <255>
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// int data types >
	var intVal int = 256
	fmt.Println(intVal)
	fmt.Printf("Variable is of type: %T \n", intVal)

	// Float datatypes
	var smallFloat float32 = 255123.4567891234 // returns total of 8 digits by trancating the first 8 digit values from whole numbers to decimal values
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var largeFloat float64 = 255123.4567891234 // Precission is high now with float64
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type: %T \n", largeFloat)

	// Default values and some aliases
	var anotherVariable int // default value for int is zero
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherVariable2 string // default value for string is empty string
	fmt.Println(anotherVariable2)
	fmt.Printf("Variable is of type: %T \n", anotherVariable2)

	// Implicit declaration of variables <Can be used to set Global variables>
	var website = "learning Golang language"
	fmt.Println(website)

	// No var style
	numberOfUsers := 500 // Not allow for variable declaration outside the function
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
