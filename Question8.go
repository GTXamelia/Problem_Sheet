package main

import (
	"fmt"
	"sort"
)

func main() {
	// Two int arrays with mixed up integer values from 1 - 6
	unsorted1 := []int{1,4,6}
	unsorted2 := []int{2,3,5}

	// This array takes the values of both unsorted interger arrays and combines them into one array
	// This array is unsorted and looks like this "[1,4,6,2,3,5]"
	formatted := append(unsorted1, unsorted2...)

	// This sorts the formatted array from lowest to greatest
	// The formatted array now looks like "[1,2,3,4,5,6]"
	sort.Ints(formatted)

	// This prints out to the user the results of the sorted array
	fmt.Print("\nMerged and sorted: ", formatted)
	
}