package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"gitee.com/3281328128/disk/internal/clean"
	"gitee.com/3281328128/disk/internal/scan"
	"gitee.com/3281328128/disk/pkg/constants"
)

func delExistsFile(filePath string) {
	outPutFile, err := os.Stat(filePath)
	if err == nil {
		if !outPutFile.IsDir() {
			err := os.Remove(filePath)
			if err != nil {
				fmt.Printf("删除文件:%s", filePath)
			}
		}
	}
}

func scanFile() {
	var rootPath string
	var outputPath string
	var thresholdByte int64

	fs := flag.NewFlagSet("start", flag.ExitOnError)
	fs.StringVar(&rootPath, "r", "", "扫描根路径")
	fs.StringVar(&outputPath, "o", "", "扫描输出文件")
	fs.Int64Var(&thresholdByte, "t", 1024*1024*1024, "文件大小阈值")
	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Println("命令行参数解析失败")
		return
	}

	if rootPath == "" {
		fmt.Println("扫描根路径不能为空")
		return
	}
	if outputPath == "" {
		fmt.Println("扫描输出文件不能为空")
		return
	}

	rootFile, err := os.Stat(rootPath)
	if err != nil {
		fmt.Println("输入文件路径错误")
		return
	}
	if !rootFile.IsDir() {
		fmt.Println("输入路径不是文件夹")
		return
	}

	outPutFile, err := os.Stat(outputPath)
	if err == nil {
		if outPutFile.IsDir() {
			outputPath = outputPath + string(os.PathSeparator) + "scan.txt"
		}
		delExistsFile(outputPath)
	}

	sacnOutFile, err := os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开扫描结果文件失败")
		return
	}
	defer sacnOutFile.Close()
	scan.Output_File = sacnOutFile
	scan.Threshold_Byte = thresholdByte

	fmt.Printf("扫描根路径:%s,扫描输出文件:%s,阈值:%d\n", rootPath, outputPath, thresholdByte)

	t := time.Now()
	size := scan.ScanFileByPath(rootPath)
	if size >= scan.Threshold_Byte {
		content := fmt.Sprintf("[%s]:%s,%d\n", constants.FILE_DIR, rootPath, size)
		fmt.Print(content)
		scan.WriteThresholdPathInfo(content)
	}
	cost_time := time.Since(t)
	fmt.Println("常规扫描消耗时间:", cost_time)
}

func main() {
	// scanFile()

	clean.CleanFile("scan.txt")

	// t_routine := time.Now()
	// fileSizeChan := make(chan int64)
	// go scan.ScanFileByPathGoRoutine(rootPath, fileSizeChan)
	// fileSumSize := <-fileSizeChan
	// fmt.Printf("文件夹:%s 包含文件大小为:%d \n", rootPath, fileSumSize)
	// routine_cost_time := time.Since(t_routine)
	// fmt.Println("Go协程扫描时间:", routine_cost_time)
}
