package main

import (
	"fmt"
	"reflect"
)

func main() {
	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	fmt.Println(reflect.DeepEqual(sm1, sm2))
}
