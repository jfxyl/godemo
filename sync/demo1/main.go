package main

import (
	"fmt"
	"sync"
)

var (
	once    sync.Once
	mapping map[string]int
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("cindy的年龄是：%d\n", get("cindy"))
	}
}

func get(name string) int {
	if mapping == nil {
		once.Do(initMapping)
	}
	return mapping[name]
}

func initMapping() {
	fmt.Println("initMapping")
	mapping = map[string]int{
		"alice": 18,
		"bob":   19,
		"cindy": 17,
	}
}
