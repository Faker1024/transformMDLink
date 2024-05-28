package file

import "strings"

// MdFile md文件
type MdFile struct {
	Path    string
	content []string
}

func (f MdFile) parseFile() {

}

func (f MdFile) save() {

}

func (f MdFile) Load() (File, error) {
	//	检查是否是md文件
	//	md文件处理逻辑
	isMD := strings.HasSuffix(f.Path, ".md")
	if !isMD {
		//	爆出错误

	} else {

	}
	return nil, nil
}
