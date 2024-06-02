package file

// MdFile md文件
type MdFile struct {
	Path    string
	content []string
}

func (f MdFile) ReplaceUrl() (File, error) {
	//TODO implement me
	panic("implement me")
}

func (f MdFile) ParseFile() (File, error) {
	//TODO implement me
	panic("implement me")
}

func (f MdFile) Save() (File, error) {
	//TODO implement me
	panic("implement me")
}

func (f MdFile) Load() (File, error) {

	return nil, nil
}
