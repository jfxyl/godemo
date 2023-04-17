package common

import (
	"os"
)

var (
	DirPath  = "./csv/tmp/"
	FilePath = DirPath + "test.csv"
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
			panic(err)
		}
	}
}
