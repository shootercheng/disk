package locales

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

//go:embed *.json
var localeFiles embed.FS

var localeMap map[string]string

var SUPPORT_LANGUAGE []string

func init() {
	getSupportLang()
}

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
	val, ok := localeMap[key]
	if !ok {
		fmt.Printf("unknown config key:%s", key)
	}
	return val
}

func getSupportLang() {
	files, err := localeFiles.ReadDir(".")
	if err != nil {
		panic("read dir error")
	}
	for _, file := range files {
		lang, found := strings.CutSuffix(file.Name(), ".json")
		if found {
			SUPPORT_LANGUAGE = append(SUPPORT_LANGUAGE, lang)
		}
	}
}
