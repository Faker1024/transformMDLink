package file

// Dir 文件夹
type Dir struct {
	Path  string
	files []File
}

func (f Dir) Load() (File, error) {
	//	文件夹处理逻辑

	return nil, nil
}
