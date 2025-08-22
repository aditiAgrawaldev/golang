package main

import "fmt"

func main() {
	fmt.Println("If else")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else {
		result = "Admin"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Less than 10")
	} else {
		fmt.Println("Greater than 10")
	}

}
