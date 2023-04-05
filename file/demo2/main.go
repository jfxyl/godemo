package main

import (
	"fmt"
	common "godemo/file"
	"os"
	"path/filepath"
	"time"
)

type node struct {
	Name  string      `json:"name"`
	IsDir bool        `json:"is_dir"`
	Type  os.FileMode `json:"type"`
	Size  int64       `json:"size"`
	//Sys      any             `json:"sys"`
	Mode     os.FileMode     `json:"mode"`
	ModTime  time.Time       `json:"mod_time"`
	SubNodes map[string]node `json:"sub_nodes"`
}

func main() {
	var (
		err  error
		tree map[string]node
	)
	//读取指定目录下的目录和文件到map
	if tree, err = readDir(common.DirPath); err != nil {
		panic(err)
	}
	fmt.Printf("%+v", tree)
}

//遍历读取目录到map
func readDir(dirPath string) (map[string]node, error) {
	var (
		err       error
		dirEntry  os.DirEntry
		dirEntrys []os.DirEntry
		fileInfo  os.FileInfo
		tree      map[string]node
		subtree   map[string]node
	)
	tree = make(map[string]node)
	if dirEntrys, err = os.ReadDir(dirPath); err != nil {
		return tree, err
	}
	for _, dirEntry = range dirEntrys {
		if dirEntry.IsDir() {
			fileInfo, err = dirEntry.Info()
			subtree, err = readDir(filepath.Join(dirPath, dirEntry.Name()))
			tree[dirEntry.Name()] = node{
				Name:     dirEntry.Name(),
				IsDir:    dirEntry.IsDir(),
				Type:     dirEntry.Type(),
				SubNodes: subtree,
			}
		} else {
			fileInfo, err = dirEntry.Info()
			tree[dirEntry.Name()] = node{
				Name:  dirEntry.Name(),
				IsDir: dirEntry.IsDir(),
				Type:  dirEntry.Type(),
				Size:  fileInfo.Size(),
				//Sys:     fileInfo.Sys(),
				Mode:    fileInfo.Mode(),
				ModTime: fileInfo.ModTime(),
			}
		}
	}
	return tree, nil
}
