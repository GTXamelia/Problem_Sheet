package main

import (
	"fmt"
	"sort"
)

func main() {
	// Unsorted int array
	unsorted := []int{5,7,9,3,4,1,8,2,6}

	// This import auto sorts the array from lowest to greatest
	sort.Ints(unsorted)

	// This then prints results to the user
	fmt.Println("Sorted: ", unsorted)

}