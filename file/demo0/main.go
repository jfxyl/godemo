package main

import (
	"bufio"
	"io/ioutil"
	"os"
)

func main() {
	var (
		err      error
		file     *os.File
		writer   *bufio.Writer
		dirPath  string
		filePath string
	)
	dirPath = "./file/tmp/testdir1/testdir2"
	filePath = dirPath + "/testfile"
	//创建一个嵌套目录
	if err = os.MkdirAll(dirPath, 0755); err != nil {
		panic(err)
	}
	//创建一个文件
	if file, err = os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755); err != nil {
		panic(err)
	}
	defer file.Close()

	//使用ioutil库写入文件
	if err = ioutil.WriteFile(filePath, []byte("ioutil.WriteFile写入测试文本\n"), 0755); err != nil { //会覆盖文件
		panic(err)
	}

	//使用os库写入文件
	if err = os.WriteFile(filePath, []byte("os.WriteFile写入测试文本\n"), 0755); err != nil { //会覆盖文件
		panic(err)
	}
	if _, err = file.Write([]byte("os.File.Write写入测试文本\n")); err != nil {
		panic(err)
	}
	if _, err = file.WriteString("os.File.WriteString写入测试文本\n"); err != nil {
		panic(err)
	}

	//使用bufio库写入文件
	writer = bufio.NewWriter(file)
	if _, err = writer.Write([]byte("bufio.Writer.Write写入测试文本\n")); err != nil {
		panic(err)
	}
	if _, err = writer.WriteString("bufio.Writer.WriteString写入测试文本\n"); err != nil {
		panic(err)
	}
	if err = writer.WriteByte('i'); err != nil {
		panic(err)
	}
	if err = writer.WriteByte('\n'); err != nil {
		panic(err)
	}
	if _, err = writer.WriteRune('我'); err != nil {
		panic(err)
	}
	if _, err = writer.WriteRune('\n'); err != nil {
		panic(err)
	}
	writer.Flush()

}
