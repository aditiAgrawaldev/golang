package main

import (
	"fmt"
	"time"
)

func cook(dish string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(dish, "step", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go cook("🍝 Pasta") // runs concurrently
	go cook("🥗 Salad") // runs concurrently

	fmt.Println("Chef is multitasking...")
	time.Sleep(5 * time.Second) // wait for goroutines
}
