package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	common "godemo/http"
	"io/ioutil"
	"net"
	"net/http"
)

func main() {
	var (
		err error
	)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var (
			err     error
			ip      string
			port    string
			reqBody []byte
			user    common.User
		)
		w.WriteHeader(http.StatusOK)
		//解析GET/POST参数
		if err = r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm error：%s\n", err.Error())
		}
		if ip, port, err = net.SplitHostPort(r.RemoteAddr); err != nil {
			fmt.Fprintf(w, "Parse host:port error：%s\n", err.Error())
		}
		//读取请求体数据
		if reqBody, err = ioutil.ReadAll(r.Body); err != nil {
			fmt.Fprintf(w, "Read Body error：%s\n", err.Error())
		}
		//重置r.Body(由于r.Body没有Seek方法，不能将指针重置到初始位置)
		r.Body = ioutil.NopCloser(bytes.NewReader(reqBody))
		//直接将请求体数据解析到struct
		if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
			fmt.Fprintf(w, "json.Decode error：%s\n", err.Error())
		}
		fmt.Println("user", user)

		fmt.Fprintf(w, "RemoteAddr：%s\n", r.RemoteAddr)
		fmt.Fprintf(w, "Host：%s\n", r.Host)
		fmt.Fprintf(w, "Ip：%s\n", ip)
		fmt.Fprintf(w, "Port：%s\n", port)
		fmt.Fprintf(w, "URL：%s\n", r.URL)
		fmt.Fprintf(w, "Method：%s\n", r.Method)
		fmt.Fprintf(w, "Form：%s\n", r.Form)
		fmt.Fprintf(w, "PostForm：%s\n", r.PostForm)
		fmt.Fprintf(w, "Body：%s\n", string(reqBody))
	})
	if err = http.ListenAndServe(":1111", nil); err != nil {
		panic(err)
	}
}
