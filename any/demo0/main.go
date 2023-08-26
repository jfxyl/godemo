package main

import (
	"cmp"
	"fmt"
)

type i int

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func add[T number](a, b T) T {
	return a + b
}

func Stringify[T fmt.Stringer](s []T) (strings []string) {
	for _, v := range s {
		strings = append(strings, v.String())
	}
	return
}

func Stringify1(s []fmt.Stringer) (strings []string) {
	for _, v := range s {
		strings = append(strings, v.String())
	}
	return
}

type stringNew struct {
	Content string
	Number  int
}

func (s stringNew) String() string {
	return s.Content
}

type People[T number, C comparable, O cmp.Ordered] struct {
	Age   T
	Sort  O
	Slice C
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 2.2))
	fmt.Println(add(i(1), i(2)))
	people := People[int, [3]int, int]{
		Age:   20,
		Sort:  1,
		Slice: [3]int{1, 2, 3},
	}
	fmt.Println(people)
	ss := []stringNew{
		stringNew{Content: "content1", Number: 1},
		stringNew{Content: "content2", Number: 2},
		stringNew{Content: "content3", Number: 3},
	}
	fmt.Println(ss)
}
