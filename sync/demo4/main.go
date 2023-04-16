package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		syncMap sync.Map
	)
	syncMap = sync.Map{}
	//写入map
	syncMap.Store("alice", 18)
	//获取key的值
	fmt.Println(syncMap.Load("alice"))
	//获取key的值,如果不存在则设置
	fmt.Println(syncMap.LoadOrStore("bob", 19))
	//遍历
	syncMap.Range(func(key, value any) bool {
		fmt.Printf("%s的年龄是%d\n", key, value)
		return true
	})
	//删除key
	syncMap.Delete("bob")
	//获取key的值并删除key
	fmt.Println(syncMap.LoadAndDelete("alice"))
}
