package main

import (
	"fmt"
	"math/rand"
	"time"
)



func main() {
	// Counter Vairable
	i := 1
	// Vairable used to hold guesses
	var num int
	var counter int
	// Vairable that calls on a function to get a random number.
	// It also sends a range to the function (1-100)
	myrand := random(1, 100)

	// vairable
	var guesses [100]int

	// Debug infromation used in development
	fmt.Println("Random Number: ", myrand)

	// For loop runs until user guesses the number correctly
    for {
		// Header showing what guess the user is on
		// This also makes the user interface much more readable seperating each guess
		fmt.Printf("\n=====================-Guess(%d)-====================", i)
		fmt.Println("\nEnter guess:")
		// This line of code takes in the user input
		n, err := fmt.Scanf("%d\n", &num)
		
		// If the user does not enter a number they will be prompted about it
        if err != nil {
            fmt.Println(n, err)
		
		
		// This statment checks to see if the user has entered a number in the valid range
		}else if num < 1 || num >= 100{
			fmt.Println("Enter next numbera number from 0 - 100")
			fmt.Println("===================================================\n")

		// This statement checks to see if the user has guesses the number incorrectly
		// It then tells the user to try again
		}else if num != myrand {
			fmt.Println("Guess Again!")

			// This checks if the user has guessed above the answer and tells them 
			if(num > myrand) {fmt.Println("You aimed too high, try a lower number this time!")}
			
			// This checks if the user has guessed below the answer and tells them 
			if(num < myrand) {fmt.Println("You aimed too low, try a higher number this time!")}
			
			fmt.Println("===================================================\n")

			// Adds incorrect guess to guess array
			guesses[num] = num

			// Adds 1 to the total guess counter
			i++

		// If either of the top statement are not true then the user has guessed correctly
		}else{
			fmt.Println("Congratulations! You guesses successfully!\n")
			fmt.Println("===================================================\n")

			// Adds incorrect guess to guess array
			guesses[num] = num

			// Adds 1 to the total guess counter
			i++

			// Exits for loop
			break
		}
	} // End of for loop
	
	// This for loop counts all entries in the array where the number !=0 
	// This is used to count only non-repeating numbers entered
	for i := 1; i < 100; i++ {
		
		if (guesses[i] != 0){counter++}

	}

	// The results section outputs stats to the user once they have won.
	fmt.Println("\n====================-(Results)-====================")
	fmt.Println(">Answer: ",myrand)
	fmt.Println(">Guesses: ",counter)
	fmt.Println(">Repeated Guesses: ",i-counter)
	fmt.Println(">Total Guesses: ",i)
	fmt.Println("===================================================\n")
}

// This function randomises a number between two numbers provided at the start of the main function
func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}