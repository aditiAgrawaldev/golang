package main

import "fmt"

// using defer keyword the function having defer keyword will execute after the completion of the function block
func main() {
	defer hello()
	fmt.Println("Start of the program")
	fmt.Println("Middle of the program")
	fmt.Println("End of the program")
}

func hello() {
	fmt.Println("Hello World")
}

// If we are having more than 1 defer functions then, they will go in stack and then LIFO manner they will execute
