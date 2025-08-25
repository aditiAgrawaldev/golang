package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (b *BankAccount) Deposit(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance += amount
	fmt.Printf("Deposited %d, Balance: %d\n", amount, b.balance)
}

func (b *BankAccount) Withdraw(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.balance >= amount {
		b.balance -= amount
		fmt.Printf("Withdrew %d, Balance: %d\n", amount, b.balance)
	} else {
		fmt.Printf("Withdraw %d failed! Balance: %d\n", amount, b.balance)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	account := &BankAccount{balance: 500}

	var wg sync.WaitGroup

	// 10 concurrent deposits & withdrawals
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			amount := rand.Intn(200) + 50
			account.Deposit(amount)
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			amount := rand.Intn(200) + 50
			account.Withdraw(amount)
		}()
	}

	wg.Wait()
	fmt.Println("Final Balance:", account.balance)
}
