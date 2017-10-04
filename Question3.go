package main

import (
	"strconv"
	"fmt"
)

func main() {
	
	// Loops 100 times
	for i := 1; i <= 100; i++ {

		// This output vairable gets re-used by each iteration of the loop
		output := "";

		// If the number is divisable by 3 then "Fizz" is added to the vairable output 
		// this is achived using the modulus with 0 remaining
		if(i % 3 == 0){output = output + "Fizz";}

		// If the number is divisable by 5 then "Buzz" is added to the vairable output 
		// this is achived using the modulus with 0 remaining
		if(i % 5 == 0){output = output + "Buzz";}

		// If both of the above conditions are met then both will output
		// the output will be "FizzBuzz" this removes redundant else if statements
		// This method allows allows for more options to be added.

		// If output is blank then the vairiable "output" = the number that the loop is in
		// This is only true if the number is not divisable by '3' or '5' such as '7'
		// If this is true then "output" becomes the number
		if(output == ""){output = strconv.Itoa(i)}

		// This line outputs the contents of the vairable "output" to the user
		// Output will be a number that is not divisable by 3/5.
		// Output will be "Fizz" if the number is diviable by 3 with no remainder
		// Output will be "Buzz" if the number is diviable by 5 with no remainder 
		// Output will be "BuzzBuzz" if the number is diviable by 3 and 5 with no remainder 
		fmt.Println(output)
		
	}
}