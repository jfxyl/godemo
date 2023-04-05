package main

import (
	"fmt"
	common "godemo/filepath"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	var (
		err     error
		abspath string
		relpath string
	)
	//相对路径和绝对路径
	fmt.Println(common.DirPath)
	if filepath.IsAbs(common.DirPath) {
		fmt.Println("绝对路径")
		abspath, _ = filepath.Abs("./")
		if relpath, err = filepath.Rel(abspath, common.DirPath); err != nil {
			panic(err)
		} else {
			fmt.Printf("其相对路径为：%s\n", relpath)
		}
	} else {
		fmt.Println("相对路径")
		if abspath, err = filepath.Abs(common.DirPath); err != nil {
			panic(err)
		} else {
			fmt.Printf("其绝对路径为：%s\n", abspath)
		}
	}
	//将path从最后一个路径分隔符后面位置分隔为两个部分
	fmt.Println(filepath.Split(common.FilePath))
	//将传入的任意数量的路径元素以路径分隔符拼接为单一路径（不同系统分隔符不同）
	fmt.Println(filepath.Join(filepath.Split(common.FilePath)))
	//将PATH或GOPATH等环境变量里的多个路径分割开（不同系统分隔符不同）
	fmt.Println(filepath.SplitList(os.Getenv("PATH")))
	//将path中的斜杠（'/'）替换为路径分隔符并返回（不同系统分隔符不同）
	fmt.Println(filepath.FromSlash(common.FilePath))
	//将path中的路径分隔符替换为斜杠（'/'）并返回（不同系统分隔符不同）
	fmt.Println(filepath.ToSlash(filepath.FromSlash(common.FilePath)))
	//返回路径中的卷名
	abspath, _ = filepath.Abs(common.DirPath)
	fmt.Println(filepath.VolumeName(abspath))
	//返回路径除去最后一个路径元素的部分
	fmt.Println(filepath.Dir(common.FilePath))
	//返回路径的最后一个元素
	fmt.Println(filepath.Base(common.FilePath))
	//返回path文件扩展名
	fmt.Println(filepath.Ext(common.FilePath))
	//通过单纯的词法操作返回和path代表同一地址的最短路径
	fmt.Println(filepath.Clean(common.FilePath))
	//返回传入的模式和路径是否符合的bool值和error
	fmt.Println(filepath.Match("*.txt", common.FilePath))
	//返回指定模式路径下能匹配到的文件路径
	fmt.Println(filepath.Glob(common.DirPath + "*.txt"))
	//对指定目录下的文件均执行WalkFunc类型方法
	filepath.Walk(common.DirPath, func(path string, info fs.FileInfo, err error) error {
		fmt.Println(path, info.Name())
		return nil
	})
}
