package main

import (
	"fmt"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type RFile struct {
}

func (*RFile) Read(p []byte) (n int, err error) {
	return len(p), nil
}

type WFile struct {
}

func (*WFile) Write(p []byte) (n int, err error) {
	return len(p), nil
}

type RWFile struct {
	RFile
	WFile
}

func main() {
	check(any(new(RFile)))
	fmt.Println("----------------------------------")
	check(any(new(WFile)))
	fmt.Println("----------------------------------")
	check(any(new(RWFile)))
	fmt.Println("----------------------------------")
	check(any(nil))
}

func check(x any) {
	//判断x是否实现了Read接口
	if y, ok := x.(Reader); !ok {
		fmt.Println("x is not Reader")
	} else {
		fmt.Println("x is Reader")
		fmt.Printf("%#v\n", y)
		y.Read([]byte{})
	}
	//判断x是否实现了Writer接口
	if y, ok := x.(Writer); !ok {
		fmt.Println("x is not Writer")
	} else {
		fmt.Println("x is Writer")
		fmt.Printf("%#v\n", y)
		y.Write([]byte{})
	}
	//判断x是否实现了ReadWriter接口
	if y, ok := x.(ReadWriter); !ok {
		fmt.Println("x is not ReadWriter")
	} else {
		fmt.Println("x is ReadWriter")
		fmt.Printf("%#v\n", y)
		y.Read([]byte{})
		y.Write([]byte{})
	}
	//判断x是否是*RFile
	if y, ok := x.(*RFile); !ok {
		fmt.Println("x is not *RFile")
	} else {
		fmt.Println("x is *RFile")
		fmt.Printf("%#v\n", y)
		y.Read([]byte{})
	}
	//判断x是否是*WFile
	if y, ok := x.(*WFile); !ok {
		fmt.Println("x is not *WFile")
	} else {
		fmt.Println("x is *WFile")
		fmt.Printf("%#v\n", y)
		y.Write([]byte{})
	}
	//判断x是否是*RWFile
	if y, ok := x.(*RWFile); !ok {
		fmt.Println("x is not *RWFile")
	} else {
		fmt.Println("x is *RWFile")
		fmt.Printf("%#v\n", y)
		y.Read([]byte{})
		y.Write([]byte{})
	}

	switch x.(type) {
	case nil:
		fmt.Println("nil")
	case Reader:
		fmt.Println("Reader")
	case Writer:
		fmt.Println("Writer")
	case ReadWriter:
		fmt.Println("ReadWriter")
	case *RFile:
		fmt.Println("*RFile")
	case *WFile:
		fmt.Println("*WFile")
	case *RWFile:
		fmt.Println("*RWFile")
	}

}
