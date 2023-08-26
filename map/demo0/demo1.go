package main

import (
	"fmt"
	"maps"
	"strings"
)

var m0 = map[string]string{"key0": "value0"}

var m1 = map[string]string{
	"key1": "value1",
	"key2": "value2",
	"key3": "value3",
	"key4": "value4",
	"key5": "value5",
	"key6": "value6",
}

func main() {
	fmt.Println(maps.Clone(m1))
	fmt.Println(maps.Equal(m1, maps.Clone(m1)))
	fmt.Println(maps.EqualFunc(m1, maps.Clone(m1), func(s string, s2 string) bool {
		return strings.ToLower(s) == strings.ToLower(s2)
	}))
	maps.Copy(m0, m1)
	fmt.Println(m0)
	maps.DeleteFunc(m1, func(s string, s2 string) bool {
		return s == "key1"
	})
	fmt.Println(m1)
}
