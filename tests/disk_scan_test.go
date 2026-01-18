package scan

import (
	"fmt"
	"testing"
	"time"

	"github.com/shootercheng/disk/internal/scan"
)

func TestScanFileByPath(t *testing.T) {
	rootPath := "E:\\code\\go\\toolset\\root"
	scan.Threshold_Byte = 5
	startTime := time.Now()
	fileSumSize := scan.ScanFileByPath(rootPath)
	costTime := time.Since(startTime)
	fmt.Println("扫描消耗时间:", costTime)
	fmt.Printf("文件夹:%s 包含文件大小为:%d \n", rootPath, fileSumSize)
}

func TestScanFileByPathGoRoutine(t *testing.T) {
	rootPath := "E:\\code\\go\\toolset\\root"
	scan.Threshold_Byte = 5
	startTime := time.Now()
	fileSizeChan := make(chan int64)
	go scan.ScanFileByPathGoRoutine(rootPath, fileSizeChan)
	fileSumSize := <-fileSizeChan
	costTime := time.Since(startTime)
	fmt.Println("协程扫描消耗时间:", costTime)
	fmt.Printf("文件夹:%s 包含文件大小为:%d \n", rootPath, fileSumSize)
}
