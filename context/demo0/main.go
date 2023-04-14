package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		c          context.Context
		cancelFunc context.CancelFunc
	)
	//创建一个5s超时和cancelFunc的context
	c, cancelFunc = context.WithTimeout(context.WithValue(context.TODO(), "name", "jfxy"), 5*time.Second)
	//启动一个协程10s后打印context携带的name值
	go run(c)
	//取一个0到1的随机数，如果大于0.5则取消context
	rand.Seed(time.Now().UnixNano())
	if rand.Float64() > 0.5 {
		cancelFunc()
	}
	select {
	//超时或者取消context都会触发
	case <-c.Done():
		if c.Err() != nil {
			if c.Err() == context.DeadlineExceeded {
				fmt.Println("timeout")
			} else if c.Err() == context.Canceled {
				fmt.Println("canceled")
			} else {
				fmt.Println("other error")
			}
		}
		fmt.Println("done")
	}
}

func run(c context.Context) {
	time.Sleep(10 * time.Second)
	//context携带的name值不会被打印
	fmt.Println(c.Value("name"))
}
