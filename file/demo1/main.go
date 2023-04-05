package main

import (
	"bufio"
	"fmt"
	common "godemo/file"
	"io"
	"io/ioutil"
	"os"
)

//golang 在1.16版本对文件读取进行了优化
func main() {
	var (
		err         error
		file        *os.File
		fileContent []byte
	)
	if file, err = os.Open(common.FilePath); err != nil {
		panic(err)
	}

	defer file.Close()

	//使用ioutil库读取文件
	if fileContent, err = ioutil.ReadFile(common.FilePath); err != nil {
		panic(err)
	}
	fmt.Println(string(fileContent))
	if fileContent, err = ioutil.ReadAll(file); err != nil {
		panic(err)
	}
	fmt.Println(string(fileContent))

	//使用os库读取文件
	if fileContent, err = os.ReadFile(common.FilePath); err != nil {
		panic(err)
	}
	fmt.Println(string(fileContent))
	if fileContent, err = readFile(common.FilePath); err != nil {
		panic(err)
	}
	fmt.Println(string(fileContent))

	//使用bufio库读取文件
	file.Seek(0, 0) //初始化读取位置
	if fileContent, err = readFile1(file); err != nil {
		panic(err)
	}
	fmt.Println(string(fileContent))
}

//读取文件
func readFile(filepath string) ([]byte, error) {
	var (
		err     error
		file    *os.File
		rlen    int
		buf     []byte
		content []byte
	)
	buf = make([]byte, 1024)
	if file, err = os.Open(filepath); err != nil {
		panic(err)
	}
	for {
		rlen, err = file.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if rlen == 0 {
			break
		}
		content = append(content, buf[:rlen]...)
	}
	return content, nil
}

//读取文件
func readFile1(file *os.File) ([]byte, error) {
	var (
		err     error
		rlen    int
		buf     []byte
		reader  *bufio.Reader
		content []byte
	)
	reader = bufio.NewReader(file)
	buf = make([]byte, 1024)
	for {
		rlen, err = reader.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if rlen == 0 {
			break
		}
		content = append(content, buf[:rlen]...)
	}
	return content, nil
}
