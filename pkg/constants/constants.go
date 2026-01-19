package constants

import (
	"github.com/shootercheng/disk/pkg/locales"
)

var (
	FILE        string
	FILE_DIR    string
	DELETE_FLAG string
)

func init() {
	FILE = locales.GetMsg("pkg_constants_001")
	FILE_DIR = locales.GetMsg("pkg_constants_002")
	DELETE_FLAG = locales.GetMsg("pkg_constants_003")
}
