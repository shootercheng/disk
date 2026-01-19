package clean

import (
	"fmt"
	"os"
	"strings"

	"github.com/shootercheng/disk/pkg/constants"
	"github.com/shootercheng/disk/pkg/locales"
)

func CleanFile(scanFilePath string) {
	fileBytes, err := os.ReadFile(scanFilePath)
	if err != nil {
		fmt.Printf(locales.GetMsg(constants.CLEAN_READ_RESULT_FAIL_KEY), scanFilePath, err.Error())
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
		fmt.Printf(locales.GetMsg(constants.CLEAN_WRITE_DELETE_RESULT_FAIL_KEY), scanFilePath, err.Error())
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
		fmt.Printf(locales.GetMsg(constants.CLEAN_CONFIRM_DELETE_KEY), delType, filePath)
		fmt.Scanln(&userInput)
		userInput = strings.ToUpper(strings.Trim(userInput, " "))
		switch userInput {
		case "Y":
			err := os.RemoveAll(filePath)
			if err != nil {
				fmt.Printf(locales.GetMsg(constants.CLEAN_DELETE_FAIL_KEY), delType, filePath, err.Error())
				indexMap[index] = err.Error()
			} else {
				fmt.Printf(locales.GetMsg(constants.CLEAN_DELETE_SUCCESS_KEY), delType, filePath)
				indexMap[index] = "ok"
			}
		case "Q":
			os.Exit(1)
		default:
			fmt.Println(locales.GetMsg(constants.CLEAN_UNKNOWN_COMMAND_KEY))
		}
	}
}
