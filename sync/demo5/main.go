package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

var (
	a int64 = 0
	b int64 = 1
	c *int64
)

func main() {
	var (
		num = 1000
		wg  = sync.WaitGroup{}
	)
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			atomic.AddInt64(&a, 1)
			//a++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("a累加%d次：%d\n", num, a)

	//b==1的话，将2赋值给b
	atomic.CompareAndSwapInt64(&b, 1, 2)
	fmt.Printf("b的新值：%d\n", b)

	atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&c)), nil, unsafe.Pointer(&b))
	fmt.Printf("c的新值：%d\n", *c)

	//将新值赋值给b，并返回b原来的旧值
	oldb := atomic.SwapInt64(&b, 3)
	fmt.Printf("b的新值：%d，b的新值：%d\n", oldb, b)

	//原子性的取值
	fmt.Printf("b的值：%d\n", atomic.LoadInt64(&b))

	//原子性的赋值
	atomic.StoreInt64(&b, 4)
	fmt.Printf("b的值：%d\n", b)
}
