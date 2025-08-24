package main

import "fmt"

type readOp struct {
	key  string
	resp chan int
}
type writeOp struct {
	key  string
	val  int
	resp chan bool
}

func main() {
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Stateful goroutine
	go func() {
		state := make(map[string]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Write and read
	resp := make(chan bool)
	writes <- writeOp{"foo", 42, resp}
	<-resp

	readResp := make(chan int)
	reads <- readOp{"foo", readResp}
	fmt.Println("Read foo:", <-readResp)
}
