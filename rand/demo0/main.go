package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		r *rand.Rand
		s []string = []string{
			"alice",
			"bob",
			"cindy",
		}
	)
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Seed(time.Now().UnixNano())
	//返回一个非负的伪随机int值
	fmt.Println(r.Int())
	//返回一个int32类型的非负的31位伪随机数
	fmt.Println(r.Int31())
	//返回一个int64类型的非负的63位伪随机数
	fmt.Println(r.Int63())
	//返回一个uint32类型的非负的32位伪随机数
	fmt.Println(r.Uint32())
	//返回一个uint64类型的非负的64位伪随机数
	fmt.Println(r.Uint64())
	//返回一个[0,n)的伪随机int值
	fmt.Println(r.Intn(10))
	//返回一个[0,n)的伪随机int32值
	fmt.Println(r.Int31n(10))
	//返回一个[0,n)的伪随机int64值
	fmt.Println(r.Int63n(10))
	//返回一个取值范围在[0.0, 1.0)的伪随机float32值
	fmt.Println(r.Float32())
	//返回一个取值范围在[0.0, 1.0)的伪随机float64值
	fmt.Println(r.Float64())
	//返回一个有n个元素的，[0,n)范围内整数的伪随机排列的切片
	fmt.Println(r.Perm(10))

	fmt.Println("他的名字是：", s[r.Intn(len(s))])
}
