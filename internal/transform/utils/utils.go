package utils

import (
	"fmt"
	"os"
	"path/filepath"
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

// GetAllMDFiles 返回指定路径下所有文件路径的字符串数组
func GetAllMDFiles(path string) ([]string, error) {
	var files []string
	// 使用 filepath.Walk 遍历目录
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 检查是否是文件
		if !info.IsDir() {
			if filepath.Ext(path) == ".md" {
				files = append(files, path)
			}
		}
		return nil
	})
	return files, err
}
