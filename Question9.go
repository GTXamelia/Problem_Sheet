package main

import (
	"fmt"
)

func main() {
	z := 6.0
	x := 2.0
		z_next := z - ((z*z - x) / (2 * z))
		fmt.Println(z_next)

}
