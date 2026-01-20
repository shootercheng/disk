package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/shootercheng/disk/internal/clean"
	"github.com/shootercheng/disk/internal/scan"
	"github.com/shootercheng/disk/pkg/constants"
	"github.com/shootercheng/disk/pkg/locales"
	"github.com/shootercheng/disk/pkg/system"
)

var (
	methodType    string
	rootPath      string
	outputPath    string
	thresholdByte int64
	language      string
)

func init() {
	flag.StringVar(&methodType, "m", "scan", "Execution method: scan or clean")
	flag.StringVar(&rootPath, "r", "", "Scan root path")
	flag.StringVar(&outputPath, "o", "scan.txt", "Scan output file")
	flag.Int64Var(&thresholdByte, "t", 1024*1024*1024, "File size threshold")
	flag.StringVar(&language, "l", "en", "Language environment")
}

func initLanguage() {
	constants.FILE = locales.GetMsg(constants.FILE_TYPE_KEY)
	constants.FILE_DIR = locales.GetMsg(constants.FILE_DIR_KEY)
	constants.DELETE_FLAG = locales.GetMsg(constants.DELETE_FLAG_KEY)
}

func loadLanguage(inputLanguage *string) {
	setLang := "en"
	for _, config := range constants.SUPPORT_LANGUAGE {
		if config == *inputLanguage {
			setLang = config
			break
		}
	}
	locales.LoadLocales(setLang)
}

func delExistsFile(filePath string) {
	outPutFile, err := os.Stat(filePath)
	if err == nil {
		if !outPutFile.IsDir() {
			err := os.Remove(filePath)
			if err != nil {
				fmt.Printf(locales.GetMsg(constants.MAIN_DELETE_FILE_FAIL_KEY), filePath, err.Error())
			} else {
				fmt.Printf(locales.GetMsg(constants.MAIN_DELETE_SCAN_FILE_SUCCESS_KEY), filePath)
			}
		}
	}
}

func scanFile() {
	if rootPath == "" {
		fmt.Println(locales.GetMsg(constants.MAIN_ROOT_PATH_EMPTY_KEY))
		os.Exit(0)
	}
	if outputPath == "" {
		fmt.Println(locales.GetMsg(constants.MAIN_OUTPUT_FILE_EMPTY_KEY))
		os.Exit(0)
	}

	rootFile, err := os.Stat(rootPath)
	if err != nil {
		fmt.Println(locales.GetMsg(constants.MAIN_INPUT_PATH_ERROR_KEY))
		os.Exit(0)
	}
	if !rootFile.IsDir() {
		fmt.Println(locales.GetMsg(constants.MAIN_INPUT_PATH_NOT_DIR_KEY))
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
		fmt.Println(locales.GetMsg(constants.MAIN_OPEN_SCAN_FILE_FAIL_KEY))
		return
	}
	defer sacnOutFile.Close()
	scan.Output_File = sacnOutFile
	scan.Threshold_Byte = thresholdByte

	fmt.Printf(locales.GetMsg(constants.MAIN_SCAN_INFO_KEY), rootPath, outputPath, thresholdByte)
	system.PrintSysInfo()

	t := time.Now()
	size := scan.ScanFileByPath(rootPath)
	if size >= scan.Threshold_Byte {
		content := fmt.Sprintf("[%s]:%s,%d\n", constants.FILE_DIR, rootPath, size)
		fmt.Print(content)
		scan.WriteThresholdPathInfo(content)
	}
	cost_time := time.Since(t)
	fmt.Printf("%s%s\n", locales.GetMsg(constants.MAIN_SCAN_TIME_KEY), cost_time)
}

func cleanFile() {
	if outputPath == "" {
		fmt.Println(locales.GetMsg(constants.MAIN_OUTPUT_FILE_PARAM_EMPTY_KEY))
		os.Exit(0)
	}
	fmt.Printf(locales.GetMsg(constants.MAIN_START_CLEAN_KEY), outputPath)
	clean.CleanFile(outputPath)
}

func main() {
	flag.Parse()
	loadLanguage(&language)
	initLanguage()

	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, locales.GetMsg(constants.MAIN_ENTER_PARAMS_KEY))
		flag.PrintDefaults()
		os.Exit(0)
	}

	switch methodType {
	case "scan":
		scanFile()
	case "clean":
		cleanFile()
	default:
		fmt.Println(locales.GetMsg(constants.MAIN_METHOD_ERROR_KEY))
		os.Exit(0)
	}
}
