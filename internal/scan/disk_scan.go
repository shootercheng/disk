package scan

import (
	"fmt"
	"os"
	"sync"

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
		fmt.Printf(locales.GetMsg("internal_scan_001"), path, err.Error())
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
			fmt.Printf(locales.GetMsg("internal_scan_002"), filePath, err.Error())
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
			fmt.Printf(locales.GetMsg("internal_scan_003"), content)
		}
	}
}

var lock sync.Mutex

func ScanFileByPathGoRoutine(path string, fileSizeChan chan int64) {
	res, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf(locales.GetMsg("internal_scan_001"), path, err.Error())
		fileSizeChan <- 0
		return
	}

	var fileSumSize int64 = 0
	for _, file := range res {
		if file.IsDir() {
			currenDirPath := path + system.FileSeparator + file.Name()
			fileSizeChan := make(chan int64)
			go ScanFileByPathGoRoutine(currenDirPath, fileSizeChan)
			dirSize := <-fileSizeChan
			if dirSize >= Threshold_Byte {
				fmt.Printf("[%s]:%s,%d\n", constants.FILE_DIR, currenDirPath, dirSize)
			}
			lock.Lock()
			fileSumSize += dirSize
			lock.Unlock()
			continue
		}
		filePath := path + system.FileSeparator + file.Name()
		fileInfo, err := file.Info()
		if err != nil {
			fmt.Printf(locales.GetMsg("internal_scan_002"), filePath, err.Error())
		} else {
			fileSize := fileInfo.Size()
			if fileSize > Threshold_Byte {
				fmt.Printf("[%s]:%s,%d\n", constants.FILE, filePath, fileSize)
			}
			lock.Lock()
			fileSumSize += fileSize
			lock.Unlock()
		}
	}
	fileSizeChan <- fileSumSize
}
