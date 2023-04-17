package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
	//s string = "aaa\nbbb"
	)
	//返回字符串表示的bool值
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("TRUE"))

	//返回字符串表示的整数值，接受正负号
	fmt.Println(strconv.ParseInt("1", 10, 64))
	fmt.Println(strconv.ParseInt("-1", 10, 64))
	fmt.Println(strconv.ParseInt("100F", 16, 64))
	fmt.Println(strconv.ParseInt("0x100F", 0, 64)) //带有前缀的base需要为0，会从字符串前置判断
	fmt.Println(strconv.Atoi("1"))

	//返回字符串表示的无符号整数值，不接受正负号
	fmt.Println(strconv.ParseUint("1", 10, 64))

	//返回字符串表示的浮点数
	fmt.Println(strconv.ParseFloat("6.1", 64))
	fmt.Println(strconv.ParseFloat("-6.1", 64))

	//返回"true"或"false"
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatBool(false))

	//返回数字指定进制的字符串表示
	fmt.Println(strconv.FormatInt(-10086, 10))
	fmt.Println(strconv.FormatUint(10086, 2))
	fmt.Println(strconv.Itoa(1))

	//返回浮点数指定格式指定位数的字符串表示
	fmt.Println(strconv.FormatFloat(6.1, 'f', 2, 64))
	fmt.Println(trimZero(6.1234560000000000000000))

	fmt.Println(string(strconv.AppendBool([]byte{}, true)))
	fmt.Println(string(strconv.AppendInt([]byte{}, -10086, 10)))
	fmt.Println(string(strconv.AppendUint([]byte{}, 10086, 10)))
	fmt.Println(string(strconv.AppendFloat([]byte{}, 6.1, 'f', 2, 64)))

}

//去掉浮点数的后置0
func trimZero(f float64) float64 {
	str := strconv.FormatFloat(f, 'f', -1, 64)
	//prec参数为-1，理论上不需要TrimRight了
	str = strings.TrimRight(str, "0")
	f, _ = strconv.ParseFloat(str, 64)
	return f
}
