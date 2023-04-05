package common

import (
	"fmt"
	"os"
)

var (
	DirPath  = "./filepath/tmp/"
	FilePath = DirPath + "test.txt"
)

func init() {
	var (
		err error
	)
	if _, err = os.Stat(DirPath); os.IsNotExist(err) {
		os.MkdirAll(DirPath, 0755)
	}
	if _, err = os.Stat(FilePath); os.IsNotExist(err) {
		if _, err = os.OpenFile(FilePath, os.O_CREATE, 0755); err != nil {
			fmt.Println(err)
		}
	}
}
