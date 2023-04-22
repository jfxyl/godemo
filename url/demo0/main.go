package main

import (
	"fmt"
	"net/url"
)

func main() {
	var (
		//err error
		s string = "https://user:password@www.baidu.com:8080/a/b?search=golang#top"
		u *url.URL
		v url.Values
	)
	//对s进行转码使之可以安全的用在URL查询里
	fmt.Println(url.QueryEscape(s))
	//用于将QueryEscape转码的字符串还原
	fmt.Println(url.QueryUnescape(url.QueryEscape(s)))
	//将url解析为URL结构体
	u, _ = url.ParseRequestURI(s) //ParseRequestURI会将#号及后面的数据放到RawQuery中,不推荐
	fmt.Printf("%+v\n", *u)
	u, _ = url.Parse(s)
	fmt.Printf("%+v\n", *u) //推荐
	//判断url是否为绝对路径
	fmt.Printf("url是否为绝对路径：%t\n", u.IsAbs())
	//获取url的请求参数
	fmt.Printf("url的请求参数：%+v\n", u.Query())
	fmt.Printf("url的请求参数：%s\n", u.RequestURI())
	fmt.Printf("url包含的用户验证信息：%s\n", u.User)

	//返回一个用户名设置为username的不设置密码的*Userinfo
	fmt.Println(url.User("username"))
	//返回一个用户名设置为username、密码设置为password的*Userinfo
	fmt.Println(url.UserPassword("username", "password"))

	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())

	//解析一个URL编码的查询字符串，并返回可以表示该查询的Values类型的字典
	v, _ = url.ParseQuery(u.RawQuery)
	fmt.Printf("url的请求参数：%+v\n", v)
	//将v编码为URL编码的查询字符串
	v = url.Values{}
	v.Set("a", "1")
	v.Set("b", "2")
	v.Add("b", "3")
	v.Add("c", "4")
	v.Del("c")
	fmt.Println(v, v.Encode())
}
