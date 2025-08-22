package main

import "fmt"

func main() {
	//done := make(chan bool)
	//
	//go func() {
	//	<-done
	//	fmt.Println("Worker 1 finished")
	//}()
	//
	//go func() {
	//	<-done
	//	fmt.Println("Worker 2 finished")
	//}()
	//
	//done <- true // Only ONE worker unblocks

	done := make(chan bool)

	go func() {
		<-done
		fmt.Println("Worker 1 finished")
	}()

	go func() {
		<-done
		fmt.Println("Worker 2 finished")
	}()

	close(done) // ✅ Both workers unblock at once

}
