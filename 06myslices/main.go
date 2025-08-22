package main

import (
	"fmt"
)

func main() {
	// var fruitList = []string{"Apple", "Banana", "Cherry", "Date", "Elderberry"}
	// fmt.Printf("Type of fruitList is %T\n", fruitList)
	// fmt.Println("Length of fruitList is ", len(fruitList))
	// fmt.Println(fruitList)
	// fruitList = append(fruitList, "Mango")
	// fmt.Println("Length of fruitList is ", len(fruitList))
	// fmt.Println(fruitList)
	// fruitList = append(fruitList[1:])
	// fmt.Println(fruitList)
	// fruitList = append(fruitList[2:4])
	// fmt.Println(fruitList)
	// fruitList = append(fruitList[:3])
	// fmt.Println(fruitList)

	// highscore := make([]int, 4)
	// highscore[0] = 234
	// highscore[1] = 945
	// highscore[2] = 0
	// highscore[3] = 676

	// highscore = append(highscore, 678, 870)

	// fmt.Println(highscore)

	// sort.Ints(highscore)
	// fmt.Println(highscore)

	var courses = []string{"Math", "Science", "English", "History", "Geography"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
