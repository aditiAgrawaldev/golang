package main

import "fmt"

func main() {
	x := 10 // Declare an integer variable

	p := &x // p is a pointer to x, storing x's memory address

	fmt.Println("Value of x:", x)             // Output: Value of x: 10
	fmt.Println("Address of x:", &x)          // Output: Address of x: (memory address)
	fmt.Println("Value of p (address):", p)   // Output: Value of p (address): (same memory address as &x)
	fmt.Println("Value pointed to by p:", *p) // Output: Value pointed to by p: 10

	*p = 20 // Change the value at the address pointed to by p

	fmt.Println("New value of x:", x) // Output: New value of x: 20 (x is now 20)
}
