package main

import (
	"fmt"
	"sync"
)

type people struct {
	name   string
	ticket int
}

func (o *people) rob() {
	cond.L.Lock()
	for amount > 0 {
		amount--
		o.ticket++
		cond.Broadcast()
		cond.Wait()
		fmt.Printf("%s抢到锁了\n", o.name)
	}

	wg.Done()
}

var (
	lock   sync.Mutex
	cond   *sync.Cond
	wg     sync.WaitGroup
	amount int = 100
)

func main() {
	var (
		alice, bob, cindy *people
	)
	wg.Add(1)
	cond = sync.NewCond(&lock)
	cond.L.Lock()
	alice = &people{name: "alice"}
	bob = &people{name: "bob"}
	cindy = &people{name: "cindy"}
	go alice.rob()
	go bob.rob()
	go cindy.rob()
	cond.L.Unlock()
	wg.Wait()
	fmt.Printf("%s抢到了%d张票\n", alice.name, alice.ticket)
	fmt.Printf("%s抢到了%d张票\n", bob.name, bob.ticket)
	fmt.Printf("%s抢到了%d张票\n", cindy.name, cindy.ticket)
}
