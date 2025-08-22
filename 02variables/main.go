package main

import "fmt"

const LoginToken string = "1234567890" //when we name a variable with a capital letter, in package scope, in starting then it is considered as public variable, so it can be used anywhere

func main() {
	var username string = "Aditi"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var undeclaredval int
	fmt.Println(undeclaredval)
	fmt.Printf("Variable is of type: %T \n", undeclaredval)

	var undeclaredstr string
	fmt.Println(undeclaredstr)
	fmt.Printf("Variable is of type: %T \n", undeclaredstr)

	//implicit type
	var implicitVal = 10
	fmt.Println(implicitVal)
	//The lexer will decide on its own what type of value is it, but if it declared it as string, we can change to int on implicit so that the problem

	//no var style
	numberOfUsers := 1000
	fmt.Println(numberOfUsers)
	//Inside any method we are allowed to use walrus operator but not outside

}
