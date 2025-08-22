package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to my program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our program:")

	//comma ok // err ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of rating is %T", input)
}
