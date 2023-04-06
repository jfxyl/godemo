package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	common "godemo/http"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	var (
		err         error
		client      *http.Client
		req         *http.Request
		resp        *http.Response
		respContent []byte
		reqUrl      string = "http://127.0.0.1:1111?name=jfxy"
		jsonData    []byte
	)
	jsonData, _ = json.Marshal(common.User{Name: "jfxy", Age: 18, Gender: "未知"})

	//http.Get
	if resp, err = http.Get(reqUrl); err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if respContent, err = ioutil.ReadAll(resp.Body); err != nil {
		panic(err)
	}
	fmt.Printf("http.Get.Response：\n%s\n", string(respContent))

	//http.Post发送kv数据
	if resp, err = http.Post(reqUrl, "application/x-www-form-urlencoded", strings.NewReader("key=value")); err != nil {
		panic(err)
	}
	if respContent, err = ioutil.ReadAll(resp.Body); err != nil {
		panic(err)
	}
	fmt.Printf("http.Post.Response：\n%s\n", string(respContent))

	//http.Post发送json数据
	if resp, err = http.Post(reqUrl, "application/json", strings.NewReader(string(jsonData))); err != nil {
		panic(err)
	}
	if respContent, err = ioutil.ReadAll(resp.Body); err != nil {
		panic(err)
	}
	fmt.Printf("http.Post.Response：\n%s\n", string(respContent))

	//http.PostForm
	if resp, err = http.PostForm(reqUrl, url.Values{"key1": {"key1_value1", "key1_value2"}, "key2": {"key2_value1"}}); err != nil {
		panic(err)
	}
	if respContent, err = ioutil.ReadAll(resp.Body); err != nil {
		panic(err)
	}
	fmt.Printf("http.PostForm.Response：\n%s\n", string(respContent))

	//http.NewRequest
	client = http.DefaultClient
	if req, err = http.NewRequest("POST", reqUrl, bytes.NewBuffer(jsonData)); err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	if resp, err = client.Do(req); err != nil {
		panic(err)
	}
	if respContent, err = ioutil.ReadAll(resp.Body); err != nil {
		panic(err)
	}
	fmt.Printf("http.NewRequest.Response：\n%s\n", string(respContent))
}
