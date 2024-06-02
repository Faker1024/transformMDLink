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

func (f Dir) ReplaceUrl() (File, error) {
	//TODO implement me
	panic("implement me")
}

func (f Dir) Load() (File, error) {
	//	文件夹处理逻辑
	files, err := utils.GetAllMDFiles(f.Path)
	if err != nil {
		return nil, err
	}
	if f.files == nil {
		f.files = make([]File, 8)
	}
	for _, file := range files {
		mdFile := &MdFile{Path: file}
		f.files = append(f.files, mdFile)
	}
	return f, nil
}

func (f Dir) ParseFile() (File, error) {
	return nil, nil
}

func (f Dir) Save() (File, error) {

	return nil, nil
}
