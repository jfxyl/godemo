package main

import (
	"errors"
	"fmt"
)

var DivError error = errors.New("除数不能为0")

type DivErr struct {
	Arg1 int
	Arg2 int
	Msg  string
}

func (e DivErr) Error() string {
	return e.Msg
}

func div(a, b int) (float64, error) {
	if b == 0 {
		return 0, DivErr{Arg1: a, Arg2: b, Msg: DivError.Error()}
	}
	return float64(a) / float64(b), nil
}

func calc(a, b int) (float64, error) {
	res, err := div(a, b)
	if err != nil {
		return res, fmt.Errorf("calc：%w", err)
	}
	return res, err
}

func main() {
	var divCantBeZero DivErr
	_, err := calc(4, 0)

	fmt.Println("errors.Is：", errors.Is(err, DivError))
	fmt.Println("errors.As：", errors.As(err, &divCantBeZero), divCantBeZero.Arg1, divCantBeZero.Arg2)
	fmt.Println("errors.Unwrap：", errors.Unwrap(err))
	fmt.Println("errors.Join：", errors.Join(err, errors.New("error1")))
}
