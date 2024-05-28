package file

// File 文件
type File interface {
	Load() (File, error)
}
