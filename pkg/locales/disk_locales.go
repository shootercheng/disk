package locales

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
)

//go:embed *.json
var localeFiles embed.FS

var localeMap map[string]string

func LoadLocales(language string) {
	filePath := fmt.Sprintf("%s.json", language)
	data, err := localeFiles.ReadFile(filePath)
	if err != nil {
		fmt.Printf("加载资源文件:%s失败\n", filePath)
		os.Exit(0)
	}
	err = json.Unmarshal(data, &localeMap)
	if err != nil {
		fmt.Printf("请检查%s格式\n", filePath)
		os.Exit(0)
	}
}

func GetMsg(key string) string {
	return localeMap[key]
}
