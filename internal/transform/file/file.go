package file

import (
	"fmt"
	"os"
)

// File 文件
type File struct {
	path string
}

// newFile 构造函数
func newFile(path string) *File {
	return &File{path: path}
}

// isDir 是否为文件夹
func (f File) isDir() bool {
	info, err := os.Stat(f.path)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return info.IsDir()
}
