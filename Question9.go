package main

import (
	"fmt"
)

func main() {
	//Vairables to be used throught the program
	z := 1.0
	x := 0.0
	z_next := 0.0
	count := 0

	// Prompt asking the user to input a number to be processed
	fmt.Printf("\n====================-Input-====================\n")	
	fmt.Print("Enter guess: ")
	// This line of code takes in the user input
	n, err := fmt.Scanf("%f", &x)
	fmt.Printf("===============================================\n")

	// If the user does not enter a number they will be prompted about it
	if(err != nil){fmt.Println(n, err)}

	// For loop that runs until the square root has been found
	for {
	z_next = z - ((z*z - x) / (2 * z))
	if(z == z_next || count == 15){break}
	z = z_next
	count++
	}

	// Square root is output to the user with UI
	fmt.Printf("\n====================-Output-===================\n")
	fmt.Printf("Square root of '%.0f' is: %.10f ", x, z)
	fmt.Printf("\n===============================================\n")

}
