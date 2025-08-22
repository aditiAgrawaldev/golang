package main

import (
	"fmt"
	"time"
)

func makePizza(pizza chan bool) {
	fmt.Println("Make a pizza...")
	time.Sleep(2 * time.Second)
	fmt.Println("Pizza finished...")
	pizza <- true
}

func makeBurger(burger chan bool) {
	fmt.Println("Make a burger...")
	time.Sleep(2 * time.Second)
	fmt.Println("Burger finished...")
	burger <- true
}

func main() {
	pizza := make(chan bool)
	burger := make(chan bool)

	go makePizza(pizza)
	go makeBurger(burger)

	select {
	case msg := <-pizza:
		fmt.Println("Recieved", msg)
	case msg := <-burger:
		fmt.Println("Recieved", msg)
	case <-time.After(1500 * time.Millisecond):
		fmt.Println("⏰ Timeout! No food yet...")
	default:
		fmt.Println("Default statement ")
	}

}
