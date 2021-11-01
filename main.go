package main

import (
	"fmt"
	"os"
	_ "path/filepath"
	_ "strings"
)

const (
	sep        = os.PathSeparator
	notLastDir = "├───"
	notLast    = "│"
	last       = "└───"
)

type dirContent []fileType

type fileType struct {
	parent   string
	filename string
	weight   int64
	isDir    bool
	content  *dirContent
}

func (dir *dirContent) lastInDir(file *fileType) bool {
	lastFileInDir := (*dir)[len(*dir)-1]
	return lastFileInDir == *file
}

func getFileWeight(path string) int64 {
	fileInfo, _ := os.Open(path)
	fileStat, _ := fileInfo.Stat()
	return fileStat.Size()
}

func getWeightString(weight int64) string {
	if weight == 0 {
		return "empty"
	}
	return fmt.Sprint(weight) + "b"
}

func drawCatalog(dir *dirContent, printFile, lastInDir bool, indent string) (scheme string) {
	for _, readFile := range *dir {
		if readFile.isDir {
			if dir.lastInDir(&readFile) {
				scheme += fmt.Sprintf("%s%s %s\n", indent, last, readFile.filename)
				scheme += drawCatalog(readFile.content, printFile, true, indent+"\t")
			} else {
				scheme += fmt.Sprintf("%s%s %s \n", indent, notLastDir, readFile.filename)
				scheme += drawCatalog(readFile.content, printFile, false, indent+notLast+"\t")
			}
		}
		if !readFile.isDir && printFile {
			if dir.lastInDir(&readFile) {
				scheme += fmt.Sprintf("%s%s %s (%s) \n", indent, last, readFile.filename, getWeightString(readFile.weight))
			} else {
				scheme += fmt.Sprintf("%s%s %s (%s) \n", indent, notLastDir, readFile.filename, getWeightString(readFile.weight))
			}
		}
	}
	return
}

func main() {
	out := os.Stdout
	// if !(len(os.Args) == 2 || len(os.Args) == 3) {
	// 	panic("usage go run main.go . [-f]")
	// }
	path := "."
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out *os.File, path string, printFile bool) error {
	dirCont, _ := readDir(path)
	fmt.Println(drawCatalog(dirCont, printFile, false, ""))
	return nil
}

func readDir(path string) (*dirContent, error) { // читает содержимое каталогов до конца дерева
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("error on read directory info : %s", err.Error())
	}

	dirCont := make(dirContent, 0, len(dir))

	for i, dirEntry := range dir {
		dirCont = append(dirCont, fileType{filename: dirEntry.Name(), isDir: dirEntry.IsDir(), parent: path, weight: getFileWeight(path + string(sep) + dirEntry.Name())})
		if dirEntry.IsDir() {
			newDirCont, err := readDir(path + string(sep) + dirEntry.Name())
			if err != nil {
				return &dirCont, err
			}
			dirCont[i].content = newDirCont
		}
	}

	return &dirCont, nil
}
