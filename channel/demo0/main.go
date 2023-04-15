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
	wg.Add(3)
	go run0(&wg)
	go run1(&wg)
	go run2(&wg)
	wg.Wait()
	fmt.Println("main end")
}

//等待子协程执行完毕
func run0(wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		waitCh chan struct{}
	)
	waitCh = make(chan struct{}, 0)
	go func() {
		time.Sleep(3 * time.Second)
		waitCh <- struct{}{}
	}()
	<-waitCh
	fmt.Println("run0 end")
}

//限制协程数量
func run1(wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		ch chan struct{}
	)
	ch = make(chan struct{}, 3)
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
		go func(i int) {
			time.Sleep(1 * time.Second)
			fmt.Printf("第%d个协程执行完毕\n", i)
			<-ch
		}(i)
	}
	for i := 0; i < cap(ch); i++ {
		ch <- struct{}{}
	}
	fmt.Println("run1 end")
}

//多路复用
func run2(wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		ch chan struct{}
	)
	ch = make(chan struct{}, 10)
	for i := 0; i < 100; i++ {
		select {
		case ch <- struct{}{}:
			fmt.Println("写入")
		case <-ch:
			fmt.Println("读取")
		}
	}
}
