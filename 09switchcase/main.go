package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in Go Lang")

	// Seed random generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 6
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)

	// Switch case example
	switch diceNumber {
	case 1:
		fmt.Println("Number is 1")
	case 2:
		fmt.Println("Number is 2")
	case 3:
		fmt.Println("Number is 3")
	case 4:
		fmt.Println("Number is 4")
	case 5:
		fmt.Println("Number is 5")
	case 6:
		fmt.Println("Number is 6")
	default:
		fmt.Println("Invalid number")
	}
}
