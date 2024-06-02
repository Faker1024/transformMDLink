package file

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
	"transformMDLink/internal/transform/utils"
)

// MdFile md文件
type MdFile struct {
	Path    string
	content []string
}

func (f *MdFile) uploadImageByPath(oldUrl string) (string, error) {
	return oldUrl, nil
}

// ParseFile 解析文件调用replaceUrl方法替换url
func (f *MdFile) ParseFile() error {
	// 匹配Markdown图片链接的正则表达式
	re := regexp.MustCompile(`!\[.*?\]\((.*?)\)`)
	inCodeBlock := false
	for i, line := range f.content {
		trimmedLine := strings.TrimSpace(line)
		// 检查是否进入或退出围栏式代码块
		if strings.HasPrefix(trimmedLine, "```") || strings.HasPrefix(trimmedLine, "~~~") {
			inCodeBlock = !inCodeBlock
			continue
		}
		// 检查缩进式代码块（必须在处理围栏式代码块之后，因为缩进式代码块不影响围栏式代码块）
		if inCodeBlock || strings.HasPrefix(trimmedLine, "    ") || strings.HasPrefix(trimmedLine, "\t") {
			continue
		}
		// 查找所有匹配的链接
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) > 1 {
				oldUrl := match[1]
				// 调用uploadImageByPath替换链接
				newUrl, err := f.uploadImageByPath(oldUrl)
				if err != nil {
					return err
				}
				// 替换旧的URL为新的URL
				f.content[i] = strings.Replace(line, oldUrl, newUrl, 1)
			}
		}
	}
	return nil
}

func (f *MdFile) generateNewMDFileName() (string, error) {
	newFileName := strings.TrimSuffix(f.Path, ".md") + "_replace.md"
	num := 1
	for utils.FileExists(newFileName) {
		newFileName = strings.TrimSuffix(f.Path, ".md") + "_" + strconv.Itoa(num) + ".md"
		num += 1
	}
	return newFileName, nil
}

// Save 将 content 逐行保存到文件名为 xxx_replace.md 的文件中
func (f *MdFile) Save() error {
	// 生成新文件名
	newFileName, err := f.generateNewMDFileName()
	if err != nil {
		return err
	}
	// 创建新文件
	file, err := os.Create(newFileName)
	if err != nil {
		return err
	}
	defer file.Close()
	// 创建一个 bufio.Writer
	writer := bufio.NewWriter(file)
	for _, line := range f.content {
		// 将内容写入文件
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	// 刷新缓冲区
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

// Load 根据path逐行读取文件然后保存到content中
func (f *MdFile) Load() error {
	file, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f.content = append(f.content, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
