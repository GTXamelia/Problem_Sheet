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

	//var validate byte
	count := len(text)-3
	pal := 0
	fail := 0

	for i := 0; i <= len(text)-3; i++ {

		if (text[i] - text[count]) == 0 {
			pal++
		}else{
			fail++
		}
		count--

	}
	fmt.Printf("\n====================================================\n")
	if(pal == len(text)-2){fmt.Printf("Palindrome!\nThat word is a palindrome with %d out of %d matching!", pal, pal)}
	if(fail>0){fmt.Printf("That word is not a palindrome,\nthere were %d leters that did not match", fail)}
	fmt.Printf("\n====================================================\n")


}