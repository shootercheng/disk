package clean

import (
	"fmt"
	"os"
	"strings"

	"gitee.com/3281328128/disk/pkg/constants"
)

func CleanFile(scanFilePath string) {
	fileBytes, err := os.ReadFile(scanFilePath)
	if err != nil {
		fmt.Printf("读取描结果文件:%s失败", scanFilePath)
		return
	}
	fileContent := string(fileBytes)
	dataLines := strings.Split(fileContent, "\n")
	indexMap := make(map[int]string)
	deleteFIleOrDir(dataLines, indexMap, constants.FILE)
	deleteFIleOrDir(dataLines, indexMap, constants.FILE_DIR)

	for index, line := range dataLines {
		if value, ok := indexMap[index]; !ok || value != "ok" {
			continue
		}
		dataLines[index] = constants.DELETE_FLAG + line
	}
	result := strings.Join(dataLines, "\n")
	err = os.WriteFile(scanFilePath, []byte(result), 0644)
	if err != nil {
		fmt.Printf("删除结果写入文件:%s失败:%s", scanFilePath, err.Error())
	}
}

func deleteFIleOrDir(lines []string, indexMap map[int]string, delType string) {
	for index, line := range lines {
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, constants.DELETE_FLAG) {
			continue
		}
		filePrefix := fmt.Sprintf("[%s]:", delType)
		content, found := strings.CutPrefix(line, filePrefix)
		if !found {
			continue
		}
		lastIndex := strings.LastIndex(content, ",")
		if lastIndex == -1 {
			continue
		}
		filePath := content[:lastIndex]
		var userInput string
		fmt.Printf("确认删除%s:%s?Y或者N:", delType, filePath)
		fmt.Scanln(&userInput)
		if userInput == "Y" || userInput == "y" {
			err := os.RemoveAll(filePath)
			if err != nil {
				fmt.Printf("删除%s:%s,失败:%s\n", delType, filePath, err.Error())
				indexMap[index] = err.Error()
			} else {
				fmt.Printf("删除%s:%s,成功\n", delType, filePath)
				indexMap[index] = "ok"
			}
		}
	}
}
