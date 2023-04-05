package main

import (
	common "godemo/file"
	"os"
	"runtime"
	"time"
)

func main() {
	var (
		err            error
		file           *os.File
		renamefilepath string
	)
	//使用os库创建临时文件
	if file, err = os.CreateTemp(common.DirPath, "test_*.txt"); err != nil {
		panic(err)
	}
	//关闭文件句柄
	file.Close()
	//文件重命名
	renamefilepath = common.DirPath + "test_rename.txt"
	if err = os.Rename(file.Name(), renamefilepath); err != nil {
		panic(err)
	}
	//文件写入内容
	if err = os.WriteFile(renamefilepath, []byte("test"), 0777); err != nil {
		panic(err)
	}
	//文件权限调整
	if err = os.Chmod(renamefilepath, 0777); err != nil {
		panic(err)
	}
	//文件所有者调整
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		if err = os.Chown(renamefilepath, os.Getuid(), os.Getgid()); err != nil {
			panic(err)
		}
	}
	//文件时间调整
	if err = os.Chtimes(renamefilepath, time.Now().Add(3600*time.Second), time.Now().Add(3600*time.Second)); err != nil {
		panic(err)
	}
	//文件清空
	if err = os.Truncate(renamefilepath, 0); err != nil {
		panic(err)
	}
	//删除文件
	if err = os.Remove(renamefilepath); err != nil {
		panic(err)
	}
}
