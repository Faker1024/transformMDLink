package file

// File 文件
type File interface {
	Load() error
	Save() error
	ParseFile() error
}
