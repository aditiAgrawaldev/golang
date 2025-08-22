package main

import "fmt"

func main() {

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("values in channel", len(ch))
	fmt.Println(<-ch)
	ch <- 3
	fmt.Println("values in channel", len(ch))
}
