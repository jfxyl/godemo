package main

import (
	"fmt"
	"sync"
)

type publicATM struct {
	balance int
	lock    sync.RWMutex
}

func (a *publicATM) Withdraw(amount int) bool {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.balance < amount || amount <= 0 {
		return false
	}
	a.balance -= amount
	return true
}

func (a *publicATM) Balance() int {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.balance
}

func main() {
	var (
		atm *publicATM
		wg  sync.WaitGroup
	)
	atm = &publicATM{
		balance: 1000,
		lock:    sync.RWMutex{},
	}
	fmt.Printf("atm现在有%d元\n", atm.Balance())
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer func() {
				wg.Done()
			}()
			if atm.Withdraw(100) {
				fmt.Println("取款100元")
			}
		}()
	}
	wg.Wait()
	fmt.Printf("atm现在有%d元\n", atm.Balance())
}
