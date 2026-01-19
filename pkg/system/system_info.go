package system

import (
	"fmt"
	"os"
	"runtime"

	"github.com/shootercheng/disk/pkg/locales"
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
	fmt.Printf(locales.GetMsg("pkg_system_001"), OsType, FileSeparator)
}
