package main

import (
	"fmt"
	"os"
)

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
			fmt.Printf("文件夹:%s 包含文件大小为:%d \n", currenDirPath, dirSize)
			fileSumSize += dirSize
			continue
		}
		filePath := path + "\\" + file.Name()
		fmt.Println(filePath)
		fileInfo, err := file.Info()
		if err != nil {
			fmt.Printf("获取文件信息:%s失败:%s\n", filePath, err.Error())
		} else {
			fileSumSize += fileInfo.Size()
		}
	}
	return fileSumSize
}

func main() {
	rootPath := "C:\\Users\\scd"
	size := ScanFileByPath(rootPath)
	fmt.Printf("文件夹%s包含文件大小为:%d \n", rootPath, size)
}
