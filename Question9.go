package main

import (
	"fmt"
)

func main() {
	z := 2.0
	x := 1.0
	z_next := 0.0

	for i := 1; i <= 10; i++ {
	z_next = z - ((z*z - x) / (2 * z))
	fmt.Print("Test ",z_next)
	x = z
	}

}
