package main

import "fmt"

func main() {
	defer fmt.Println("Deferred call executed even after panic")
	panic("Something went wrong!") // crash
}
