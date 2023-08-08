package main

import (
	"bufio"
	"fmt"
	"os"
)

// Take inputs from user
func main() {
	topic := "User inputs"
	fmt.Println(topic)

	// Reader read data from OS
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok || err
	input, _ := reader.ReadString('\n') // retuns string
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T, \n ", input)
}

//  We need to convert the returned string into appropriate data type
