package system

import (
	"fmt"
	"os"
	"runtime"
)

var (
	OsType        string
	FileSeparator string
)

func init() {
	OsType = runtime.GOOS
	FileSeparator = string(os.PathSeparator)
}

func PrintSysInfo() {
	fmt.Printf("当前系统类型:%s,文件路径分隔符:%s\n", OsType, FileSeparator)
}
