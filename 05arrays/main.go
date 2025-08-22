package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Cherry"

	fmt.Println("Fruit list is", fruitList)
	fmt.Println("Length of Fruit list is", len(fruitList))

	var vegList = [10]string{"potato", "beans", "tomato"}
	fmt.Println("Veg list is", vegList)
	fmt.Println("Length of Veg list is", len(vegList))

}
