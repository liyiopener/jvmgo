package main

import (
	"os"
)

//路径分隔符
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {

	//寻找和加载 class 文件
	readClass(className string) ([]byte, Entry, error)

	// toString 方法
	String() string
}
