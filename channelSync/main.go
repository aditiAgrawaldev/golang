package main

import (
	"fmt"
	"time"
)

func cook(done chan bool) {
	fmt.Println("Cooking started...")
	time.Sleep(2 * time.Second) // Making food
	fmt.Println("Cooking finished...")
	done <- true
}

func main() {
	done := make(chan bool)
	go cook(done)
	<-done
	fmt.Println("Finally Eat...")
}

// /Without Channel sync
package main

import (
	"fmt"
	"time"
)

func cook() {
	fmt.Println("cooking...")
	time.Sleep(2 * time.Second)
	fmt.Println("done cooking..")
}

func main() {
	go cook()
	time.Sleep(1 * time.Second)
	fmt.Println("Finally Eat....")
}
