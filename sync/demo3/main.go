package main

import (
	"fmt"
	"sync"
)

type people struct {
	Name string
	Age  int
}

func main() {
	var (
		pool  sync.Pool
		alice people
	)
	pool = sync.Pool{New: func() any {
		return people{}
	}}
	//取出一个people,没有的话，则使用New方法创建一个
	alice = pool.Get().(people)
	alice.Name = "alice"
	alice.Age = 18
	//将people放入池中
	pool.Put(alice)
	//从池中取出people
	fmt.Println(pool.Get().(people))
}
