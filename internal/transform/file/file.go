package file

// File 文件
type File interface {
	Load() (File, error)
	Save() (File, error)
	ParseFile() (File, error)
	ReplaceUrl() (File, error)
}
