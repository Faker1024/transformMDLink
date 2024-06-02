package transform

import (
	"transformMDLink/internal/transform/config"
	"transformMDLink/internal/transform/file"
	"transformMDLink/internal/transform/utils"
)

func Run(config config.Config) {
	var f file.File
	flag := utils.IsDir(config.Path())
	if flag {
		f = file.Dir{Path: config.Path()}
	} else {
		f = file.MdFile{Path: config.Path()}
	}
	f.Load()
}
