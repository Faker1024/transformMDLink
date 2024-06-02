package file

import (
	"fmt"
	"testing"
	"transformMDLink/internal/transform/utils"
)

//func TestRead(t *testing.T) {
//	read, err := file.Read("/home/jack/GolandProjects/transformMDLink/assets/test.md")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(read)
//}
//
//func TestIsDir(t *testing.T) {
//	dir, err := file.IsDir("/home/jack/GolandProjects/transformMDLink/assets/test.md")
//	if err != nil {
//		return
//	}
//	fmt.Println(dir)
//	dir, err = file.IsDir("/home/jack/GolandProjects/transformMDLink/assets")
//	if err != nil {
//		return
//	}
//	fmt.Println(dir)
//}
//
//func TestGetFiles(t *testing.T) {
//	files, err := file.GetMDFiles("/home/jack/GolandProjects/transformMDLink")
//	if err != nil {
//		return
//	}
//	for _, s := range files {
//		fmt.Println(s)
//	}
//}

func TestGetFirstLevelFilesAndDirs(t *testing.T) {
	dirs, err := utils.GetAllMDFiles("/home/jack/GolandProjects/transformMDLink")
	if err != nil {
		return
	}
	for _, dir := range dirs {
		fmt.Println(dir)
	}
}
