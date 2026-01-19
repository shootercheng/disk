package system

import (
	"fmt"
	"os"
	"runtime"

	"github.com/shootercheng/disk/pkg/constants"
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
	fmt.Printf(locales.GetMsg(constants.SYSTEM_INFO_MSG_KEY), OsType, FileSeparator)
}
