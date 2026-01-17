package scan

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	"gitee.com/3281328128/disk/pkg/constants"
)

var Threshold_Byte int64 = 1024 * 1024 * 1024

var Output_File *os.File

var separator string

func init() {
	osType := runtime.GOOS
	separator = string(os.PathSeparator)
	fmt.Printf("当前系统类型:%s,文件路径分隔符:%s\n", osType, separator)
}

func ScanFileByPath(path string) int64 {
	// fmt.Println("开始扫描文件夹:", path)
	res, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("读取文件夹 %s 失败:%s\n", path, err.Error())
		return 0
	}

	var fileSumSize int64 = 0
	for _, file := range res {
		if file.IsDir() {
			currenDirPath := path + separator + file.Name()
			dirSize := ScanFileByPath(currenDirPath)
			if dirSize >= Threshold_Byte {
				content := fmt.Sprintf("[%s]:%s,%d\n", constants.FILE_DIR, currenDirPath, dirSize)
				fmt.Print(content)
				WriteThresholdPathInfo(content)
			}
			fileSumSize += dirSize
			continue
		}
		filePath := path + separator + file.Name()
		fileInfo, err := file.Info()
		if err != nil {
			fmt.Printf("获取文件信息:%s失败:%s\n", filePath, err.Error())
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
			fmt.Printf("路径%s写入输出文件失败\n", content)
		}
	}
}

var lock sync.Mutex

func ScanFileByPathGoRoutine(path string, fileSizeChan chan int64) {
	// fmt.Println("开始扫描文件夹:", path)
	res, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("读取文件夹 %s 失败:%s\n", path, err.Error())
		fileSizeChan <- 0
		return
	}

	var fileSumSize int64 = 0
	for _, file := range res {
		if file.IsDir() {
			currenDirPath := path + separator + file.Name()
			fileSizeChan := make(chan int64)
			go ScanFileByPathGoRoutine(currenDirPath, fileSizeChan)
			dirSize := <-fileSizeChan
			if dirSize >= Threshold_Byte {
				fmt.Printf("[%s]:%s 文件大小为:%d \n", constants.FILE_DIR, currenDirPath, dirSize)
			}
			lock.Lock()
			fileSumSize += dirSize
			lock.Unlock()
			continue
		}
		filePath := path + separator + file.Name()
		fileInfo, err := file.Info()
		if err != nil {
			fmt.Printf("获取文件信息:%s失败:%s\n", filePath, err.Error())
		} else {
			fileSize := fileInfo.Size()
			if fileSize > Threshold_Byte {
				fmt.Printf("[%s]:%s 大小为:%d \n", constants.FILE, filePath, fileSize)
			}
			lock.Lock()
			fileSumSize += fileSize
			lock.Unlock()
		}
	}
	fileSizeChan <- fileSumSize
}
