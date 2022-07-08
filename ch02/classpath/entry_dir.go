package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// 目录形式的类路径
type DirEntry struct {
	//绝对路径
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

/*
1. 目录和 className 拼接全路径
2. io 读取文件内容
*/
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) string() string {
	return self.absDir
}
