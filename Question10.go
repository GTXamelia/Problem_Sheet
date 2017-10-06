package main

import (
	"fmt"
	"os"
	"bufio"
)

// Function that reverses the string using a rune
// The function gets its string from the main function where the user is prompted to enter a word
func reverse(text string) string {

    input := []rune(text)
	output := []rune{}
	
	// For loop starts from the top of the string gets its letter and adds it to a rune
	// This method repeats until the entire word has been put threw it and all letters added to the ouput vairable
    for i := len(input) - 1; i >= 0; i-- {
        output = append(output, input[i])
    }

	// Output is then retuned as a string to the main method
    return string(output)
}

func main() {

	// The user is prompted to enter a string to reverse
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n====================-Input-====================\n")
	fmt.Print("Enter Text: \n")
	text, _ := reader.ReadString('\n')
	fmt.Printf("===============================================\n")

	// Runs the reverse method and sends the string to it
	// The reverse variable then becomes the outcome which is the word reveresed
	reversed := reverse(text)

	// The user is then shown the results of the outcome
	fmt.Printf("\n====================-Output-===================\n")
	fmt.Printf("Text reversed: %s",reversed)
	fmt.Printf("\n===============================================\n")
}
