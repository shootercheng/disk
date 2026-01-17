package main

import (
	"fmt"
	"os"
)

var threshold_byte int64 = 5

func ScanFileByPath(path string) int64 {
	// fmt.Println("开始扫描文件夹:", path)
	res, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("读取文件夹失败")
	}

	var fileSumSize int64 = 0
	for _, file := range res {
		if file.IsDir() {
			currenDirPath := path + "\\" + file.Name()
			dirSize := ScanFileByPath(currenDirPath)
			if dirSize >= threshold_byte {
				fmt.Printf("文件夹:%s 包含文件大小为:%d \n", currenDirPath, dirSize)
			}
			fileSumSize += dirSize
			continue
		}
		filePath := path + "\\" + file.Name()
		fileInfo, err := file.Info()
		if err != nil {
			fmt.Printf("获取文件信息:%s失败:%s\n", filePath, err.Error())
		} else {
			fileSize := fileInfo.Size()
			if fileSize > threshold_byte {
				fmt.Printf("文件:%s 大小为:%d \n", filePath, fileSize)
			}
			fileSumSize += fileSize
		}
	}
	return fileSumSize
}

func main() {
	rootPath := "E:\\code\\go\\toolset\\root"
	size := ScanFileByPath(rootPath)
	fmt.Printf("文件夹%s包含文件大小为:%d \n", rootPath, size)
}
