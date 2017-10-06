package main

import (
	"fmt"
	"os"
	"bufio"
)



func reverse(text string) string {

    input := []rune(text)
	output := []rune{}
	
    for i := len(input) - 1; i >= 0; i-- {
        output = append(output, input[i])
    }
	
    return string(output)
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n====================-Input-====================\n")
	fmt.Print("Enter Text: \n")
	text, _ := reader.ReadString('\n')
	fmt.Printf("===============================================\n")

	reversed1 := reverse(text)

	fmt.Printf("\n====================-Output-===================\n")
	
	fmt.Printf("Text reversed: %s",reversed1)
	fmt.Printf("\n===============================================\n")
}
