package main

import "fmt"

func main() {
	// Create a slice using a composite literal
	numbers := []int{10, 20, 30, 40, 50, 70}
	fmt.Println("Original slice:", numbers)

	// Get length and capacity
	fmt.Printf("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))

	// Append an element
	numbers = append(numbers, 60)
	fmt.Println("After append:", numbers)

	// Create a sub-slice
	subSlice := numbers[1:4]
	fmt.Println("Sub-slice:", subSlice)

	// Modify an element in the sub-slice (affects the original slice)
	subSlice[0] = 25
	fmt.Println("Sub-slice after modification:", subSlice)
	fmt.Println("Original slice after sub-slice modification:", numbers)
}
