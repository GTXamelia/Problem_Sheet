package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	// Allows the user to input a word so it can be processed
	// Also has a small UI to seperate each step of the program
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n====================================================\n")
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Printf("====================================================\n")

	// This variable gets the size of the word
	count := len(text)-3

	// These two vairables are used as counter to determine if a word is a palindrome 
	pal := 0
	fail := 0

	// For loop goes up threw the word letter by letter and another counter comes down
	// as the loop goes up and compares each letter (1st and last - 2nd and 2nd-last) until it meets in the middle
	for i := 0; i <= len(text)-3; i++ {

		// If the last letter and first equal 0 then that part is palindrome
		// Using the ascii table we know a = 97, so Area would be palindrome as 97 - 97 = 0
		// using this method we can find if a word is palindrome and how many letters in the sentence are
		if (text[i] - text[count]) == 0 {
			pal++
		}else{
			fail++
		}
		// Sets counter to go down from top of string letter by letter and compare with loop going up
		count--

	}

	// Outputs the result of the palindrome test to the user whether it failed or passed 
	fmt.Printf("\n====================================================\n")
	if(pal == len(text)-2){fmt.Printf("Palindrome!\nThat word is a palindrome with %d out of %d matching!", pal, pal)}
	if(fail>0){fmt.Printf("That word is not a palindrome,\nthere were %d leters that did not match", fail)}
	fmt.Printf("\n====================================================\n")
}