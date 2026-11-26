package scan

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/shootercheng/disk/pkg/constants"
	"github.com/shootercheng/disk/pkg/locales"
	"github.com/shootercheng/disk/pkg/system"
)

var (
	Threshold_Byte int64 = constants.GB_CAP
	Output_File    *os.File
)

func ScanInputPath(path string) {
	size := ScanFileByPath(path)
	if size >= Threshold_Byte {
		content := fmt.Sprintf("[%s]:%s,%s\n", constants.FILE_DIR, path, ConvertUnit(size))
		fmt.Print(content)
		WriteThresholdPathInfo(content)
	}
}

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
				content := fmt.Sprintf("[%s]:%s,%s\n", constants.FILE_DIR, currenDirPath, ConvertUnit(dirSize))
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
				content := fmt.Sprintf("[%s]:%s,%s\n", constants.FILE, filePath, ConvertUnit(fileSize))
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

func ConvertUnit(inputSize int64) string {
	if inputSize < constants.KB_CAP {
		return strconv.FormatInt(inputSize, 10) + "B"
	}
	if inputSize < constants.MB_CAP {
		return DivConversion(inputSize, constants.KB_CAP, "KB")
	}
	if inputSize < constants.GB_CAP {
		return DivConversion(inputSize, constants.MB_CAP, "MB")
	}
	return DivConversion(inputSize, constants.GB_CAP, "GB")
}

func DivConversion(inputSize int64, unitSize int64, unitStr string) string {
	if inputSize%unitSize == 0 {
		return strconv.FormatInt(inputSize/unitSize, 10) + unitStr
	} else {
		result := float64(inputSize) / float64(unitSize)
		format := fmt.Sprintf("%.2f", result)
		format = strings.TrimRight(format, "0") + unitStr
		return format
	}
}
