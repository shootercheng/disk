package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"gitee.com/3281328128/disk/internal/clean"
	"gitee.com/3281328128/disk/internal/scan"
	"gitee.com/3281328128/disk/pkg/constants"
	"gitee.com/3281328128/disk/pkg/system"
)

var (
	methodType    string
	rootPath      string
	outputPath    string
	thresholdByte int64
)

func init() {
	flag.StringVar(&methodType, "m", "scan", "执行方法:scan或者clean")
	flag.StringVar(&rootPath, "r", "", "扫描根路径")
	flag.StringVar(&outputPath, "o", "scan.txt", "扫描输出文件")
	flag.Int64Var(&thresholdByte, "t", 1024*1024*1024, "文件大小阈值")
}

func delExistsFile(filePath string) {
	outPutFile, err := os.Stat(filePath)
	if err == nil {
		if !outPutFile.IsDir() {
			err := os.Remove(filePath)
			if err != nil {
				fmt.Printf("删除文件:%s,失败:%s\n", filePath, err.Error())
			} else {
				fmt.Printf("删除扫描输出文件:%s 成功\n", filePath)
			}
		}
	}
}

func scanFile() {
	if rootPath == "" {
		fmt.Println("扫描根路径不能为空")
		os.Exit(0)
	}
	if outputPath == "" {
		fmt.Println("扫描输出文件不能为空")
		os.Exit(0)
	}

	rootFile, err := os.Stat(rootPath)
	if err != nil {
		fmt.Println("输入文件路径错误")
		os.Exit(0)
	}
	if !rootFile.IsDir() {
		fmt.Println("输入路径不是文件夹")
		os.Exit(0)
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
	system.PrintSysInfo()

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

func cleanFile() {
	if outputPath == "" {
		fmt.Println("扫描输出文件参数不能为空")
		os.Exit(0)
	}
	fmt.Printf("开始执行文件清理任务,扫描结果文件路径:%s\n", outputPath)
	clean.CleanFile(outputPath)
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "请输入提示的命令行参数, 无默认值(default)的为必须输入参数")
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()

	switch methodType {
	case "scan":
		scanFile()
	case "clean":
		cleanFile()
	default:
		fmt.Println("命令行参数m错误,请输入scan或者clean")
		os.Exit(0)
	}
}
