package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func student(id int, wg *sync.WaitGroup, ratings chan int) {
	defer wg.Done()
	// Random delay to simulate response time
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	// Random rating between 1–10
	rating := rand.Intn(10) + 1
	fmt.Printf("Student %d rated: %d\n", id, rating)
	ratings <- rating
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	ratings := make(chan int, 200)

	// 200 students
	for i := 1; i <= 200; i++ {
		wg.Add(1)
		go student(i, &wg, ratings)
	}

	// Close channel after all done
	go func() {
		wg.Wait()
		close(ratings)
	}()

	// Collect ratings
	total := 0
	count := 0
	for r := range ratings {
		total += r
		count++
	}

	average := float64(total) / float64(count)
	fmt.Printf("\nAverage Rating: %.2f\n", average)
}
