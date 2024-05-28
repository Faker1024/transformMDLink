package utils

import (
	"fmt"
	"os"
)

// IsDir 是否为文件夹
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return info.IsDir()
}
