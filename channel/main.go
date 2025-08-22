package main

import (
	"fmt"
	"time"
)

func chef(order string, ch chan string) {
	fmt.Println("Chef Started cooking", order)
	time.Sleep(2 * time.Second)
	fmt.Println("Chef Finished cooking", order)
	ch <- order + "Food is ready"
}

func main() {
	orders := make(chan string)
	go chef("🍕 Pizza", orders)
	go chef("🥗 Salad", orders)
	go chef("🍝 Pasta", orders)

	// Wait for chefs to finish (3 orders)
	for i := 0; i < 3; i++ {
		ready := <-orders
		fmt.Println("🍽️ Waiter received:", ready)
	}

	fmt.Println("✅ All orders served!")
}

package main

import (
	"fmt"
	"time"
)

func chef(order string) {
	fmt.Println("👨‍🍳 Chef started cooking:", order)
	time.Sleep(2 * time.Second)
	fmt.Println("✅ Chef finished cooking:", order)
	// No way to notify main that food is ready
}

func main() {
	go chef("🍕 Pizza")
	go chef("🥗 Salad")
	go chef("🍝 Pasta")

	fmt.Println("Waiter waiting for dishes...")

	// Try waiting manually
	time.Sleep(2 * time.Second)

	fmt.Println("Waiter serves dishes (but doesn’t know if chefs finished)")
}
