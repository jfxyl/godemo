package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		wg sync.WaitGroup
	)
	wg.Add(2)
	go func() {
		run0()
		wg.Done()
	}()
	go func() {
		run1()
		wg.Done()
	}()
	wg.Wait()
}

func run0() {
	var (
		timer0, timer1 *time.Timer
	)
	timer0 = time.AfterFunc(2*time.Second, func() {
		fmt.Println("run0：timer0 do something")
	})
	timer1 = time.AfterFunc(10*time.Second, func() {
		fmt.Println("run0：timer1 do something")
	})
	select {
	case <-timer0.C:
		fmt.Println("run0：timer0 expired")
	case <-timer1.C:
		fmt.Println("run0：timer1 expired")
	case <-time.After(5 * time.Second):
		timer0.Stop()
		timer1.Stop()
		fmt.Println("run0：Timer cancelled")
	}
}

func run1() {
	var (
		timer    *time.Timer
		ticker   *time.Ticker
		stopChan <-chan time.Time
	)
	timer = time.NewTimer(1 * time.Second)
	ticker = time.NewTicker(2 * time.Second)
	stopChan = time.After(10 * time.Second)
	for {
		select {
		case <-timer.C:
			fmt.Println("run1：Timer expired")
			timer.Reset(1 * time.Second)
		case <-ticker.C:
			fmt.Println("run1：Ticker expired")
		case <-stopChan:
			fmt.Println("run1：Stop")
			goto END
		}
	}
END:
}
