package main

import (
	"fmt"
)

func main() {

	// Vairable to be used in calculating the factorial equation
	var facto int = 1;

	// This for loop runs 10 times before stopping
	// it will be used to get the factorial of 10 as it runs 10 times. 
	for i := 1; i <= 10; i++ {
		
		// facto = itself multiplied by the iteration of the loop
		// facto = (1 * 1 = 1), (1 * 2 = 2), (2 * 3 = 6), (6 * 4 = 24)..... 
		// and so on until 10 iteration have been comleted
		facto=facto*i;

		// Prints out result of the factorial equation
		fmt.Println(facto)
	}

}