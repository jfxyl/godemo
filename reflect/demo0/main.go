package main

import (
	"fmt"
	"os"
	"reflect"
)

type PersonInterface interface {
	SayHello() string
}

type Person struct {
	Name    string
	Age     int
	Address string
}

func (p Person) SayHello() string {
	return fmt.Sprintf("hello, my name is %s", p.Name)
}

func main() {
	var (
		s               any
		i               any
		w               any
		person          *Person
		personInterface PersonInterface

		//s string    = "hello"
		//i int64     = 1
		//w io.Writer = os.Stdout
	)

	s = "hello"
	i = 1
	w = os.Stdout

	fmt.Println(reflect.TypeOf(s), reflect.ValueOf(s), reflect.ValueOf(s).Interface(), reflect.ValueOf(s).String())
	fmt.Println(reflect.TypeOf(i), reflect.ValueOf(i), reflect.ValueOf(i).Interface(), reflect.ValueOf(i).String())
	fmt.Println(reflect.TypeOf(w), reflect.ValueOf(w), reflect.ValueOf(w).Interface(), reflect.ValueOf(w).String())
	fmt.Println(reflect.TypeOf(nil))
	fmt.Printf("%T%T\n", i, w)

	person = &Person{
		Name:    "zhangsan",
		Age:     18,
		Address: "beijing",
	}
	personInterface = *person
	do(person)
	do(*person)
	do(personInterface)

}

func do(param any) {
	var (
		t reflect.Type
		v reflect.Value
	)
	//获取param的类型
	t = reflect.TypeOf(param)
	fmt.Printf("param.type:%s\n", t)
	//获取param的值
	v = reflect.ValueOf(param)
	fmt.Printf("param.value: %#v\n", v)
	//获取param值的分类
	fmt.Printf("param.value.kind: %d\n", v.Kind())
	//如果param值是接口，则v.Elem()返回接口包含的值
	//如果param值是指针，则v.Elem()返回指针指向的值
	//如果param值是其他类型，则panic
	if v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	//获取param值的Elem()的分类
	fmt.Printf("param.value.elem.kind: %d\n", v.Kind())
	if v.Kind() == reflect.Struct {
		//获取param值指向的结构体的字段名和字段值
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("param.value.%s: %v\n", v.Type().Field(i).Name, v.FieldByName(v.Type().Field(i).Name))
		}
		//调用param值指向的结构体的方法
		for i := 0; i < v.NumMethod(); i++ {
			fmt.Printf("param.method.%s: %s\n", t.Method(i).Name, v.MethodByName(t.Method(i).Name).Call(nil))
		}
	}

}
