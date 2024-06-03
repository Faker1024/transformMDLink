package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

// FileExists 检查给定路径是否存在文件
func FileExists(path string) bool {
	_, err := os.Stat(path)
	// os.Stat returns an error if the file does not exist or the path is invalid
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

// ReadImage 读取图片文件，支持本地文件和网络URL
func ReadImage(path string) ([]byte, error) {
	// 检查是否为网络URL
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return readImageFromURL(path)
	}
	// 否则视为本地文件
	return readImageFromFile(path)
}

// readImageFromURL 从网络URL读取图片
func readImageFromURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch image, status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

// readImageFromFile 从本地文件读取图片
func readImageFromFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}
