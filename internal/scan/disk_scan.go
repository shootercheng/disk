package scan

import (
	"fmt"
	"os"

	"github.com/shootercheng/disk/pkg/constants"
	"github.com/shootercheng/disk/pkg/locales"
	"github.com/shootercheng/disk/pkg/system"
)

var (
	Threshold_Byte int64 = 1024 * 1024 * 1024
	Output_File    *os.File
)

func ScanFileByPath(path string) int64 {
	res, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf(locales.GetMsg(constants.SCAN_READ_FOLDER_FAIL_KEY), path, err.Error())
		return 0
	}

	var fileSumSize int64 = 0
	for _, file := range res {
		if file.IsDir() {
			currenDirPath := path + system.FileSeparator + file.Name()
			dirSize := ScanFileByPath(currenDirPath)
			if dirSize >= Threshold_Byte {
				content := fmt.Sprintf("[%s]:%s,%d\n", constants.FILE_DIR, currenDirPath, dirSize)
				fmt.Print(content)
				WriteThresholdPathInfo(content)
			}
			fileSumSize += dirSize
			continue
		}
		filePath := path + system.FileSeparator + file.Name()
		fileInfo, err := file.Info()
		if err != nil {
			fmt.Printf(locales.GetMsg(constants.SCAN_GET_FILE_INFO_FAIL_KEY), filePath, err.Error())
		} else {
			fileSize := fileInfo.Size()
			fileSumSize += fileSize
			if fileSize >= Threshold_Byte {
				content := fmt.Sprintf("[%s]:%s,%d\n", constants.FILE, filePath, fileSize)
				fmt.Print(content)
				WriteThresholdPathInfo(content)
			}
		}
	}
	return fileSumSize
}

func WriteThresholdPathInfo(content string) {
	if Output_File != nil {
		_, err := Output_File.WriteString(content)
		if err != nil {
			fmt.Printf(locales.GetMsg(constants.SCAN_WRITE_PATH_FAIL_KEY), content)
		}
	}
}
