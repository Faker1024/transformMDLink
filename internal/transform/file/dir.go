package file

import (
	"transformMDLink/internal/transform/utils"
)

// Dir 文件夹
type Dir struct {
	Path  string
	files []File
	error error
}

func (f Dir) Load() error {
	//	文件夹处理逻辑
	files, err := utils.GetAllMDFiles(f.Path)
	if err != nil {
		return err
	}
	if f.files == nil {
		f.files = make([]File, 8)
	}
	for _, file := range files {
		mdFile := &MdFile{Path: file}
		f.files = append(f.files, mdFile)
	}
	return nil
}

func (f Dir) ParseFile() error {
	return nil
}

func (f Dir) Save() error {

	return nil
}
