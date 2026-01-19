package locales

import (
	"testing"

	"github.com/shootercheng/disk/pkg/locales"
)

func TestLoadLocalesZh(t *testing.T) {
	locales.LoadLocales("zh")
	key := "pkg_system_001"
	val := locales.GetMsg(key)
	if val != "当前系统类型:%s,文件路径分隔符:%s\n" {
		t.Errorf("加载国际化key %s 失败", key)
	}
}

func TestLoadLocaleEn(t *testing.T) {
	locales.LoadLocales("en")
	key := "pkg_system_001"
	val := locales.GetMsg(key)
	if val != "Current system type: %s,file path separator: %s\n" {
		t.Errorf("加载国际化key %s 失败", key)
	}
}
