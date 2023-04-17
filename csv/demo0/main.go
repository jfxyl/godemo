package main

import (
	"encoding/csv"
	"fmt"
	common "godemo/csv"
	"os"
	"strconv"
)

type People struct {
	Name   string
	Age    int64
	Gender string
}

func main() {
	write()
	read()
}

func write() {
	var (
		err     error
		file    *os.File
		writer  *csv.Writer
		header  []string
		peoples []People
	)

	if file, err = os.Create(common.FilePath); err != nil {
		panic(err)
	}
	defer file.Close()

	writer = csv.NewWriter(file)
	defer writer.Flush()

	header = []string{"姓名", "年龄", "性别"}
	peoples = []People{
		{"alice", 18, "女"},
		{"bob", 19, "男"},
		{"cindy", 19, "女"},
	}

	if err = writer.Write(header); err != nil {
		panic(err)
	}
	for _, people := range peoples {
		record := []string{people.Name, strconv.FormatInt(people.Age, 10), people.Gender}
		if err = writer.Write(record); err != nil {
			panic(err)
		}
	}
}

func read() {
	var (
		err     error
		file    *os.File
		reader  *csv.Reader
		peoples []People
		records [][]string
	)

	if file, err = os.Open(common.FilePath); err != nil {
		panic(err)
	}
	defer file.Close()
	reader = csv.NewReader(file)
	if records, err = reader.ReadAll(); err != nil {
		panic(err)
	}
	for i, item := range records {
		//下面有直接取下标的操作，必须验证
		if len(item) < 3 {
			panic("csv 文件不正确")
		}
		if i == 0 {
			continue
		}
		age, err := strconv.ParseInt(item[1], 10, 64)
		if err != nil {
			panic(err)
		}
		peoples = append(peoples, People{
			Name: item[0], Age: age, Gender: item[2],
		})
	}
	fmt.Println(peoples)
}
