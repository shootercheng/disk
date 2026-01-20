package locales

import (
	"testing"

	"github.com/shootercheng/disk/pkg/constants"
	"github.com/shootercheng/disk/pkg/locales"
)

func TestLoadLocalesZh(t *testing.T) {
	locales.LoadLocales("zh")
	key := constants.SYSTEM_INFO_MSG_KEY
	val := locales.GetMsg(key)
	if val != "当前系统类型:%s,文件路径分隔符:%s\n" {
		t.Errorf("加载国际化key %s 失败", key)
	}
}

func TestLoadLocaleEn(t *testing.T) {
	locales.LoadLocales("en")
	key := constants.SYSTEM_INFO_MSG_KEY
	val := locales.GetMsg(key)
	if val != "Current system type:%s,file path separator:%s\n" {
		t.Errorf("加载国际化key %s 失败", key)
	}
}

func TestAllLocaleKeys(t *testing.T) {
	// All available keys in the locale files
	keys := []string{
		"file_type",
		"file_dir",
		"delete_flag",
		"system_info_msg",
		"scan_read_folder_fail",
		"scan_get_file_info_fail",
		"scan_write_path_fail",
		"clean_read_result_fail",
		"clean_write_delete_result_fail",
		"clean_confirm_delete",
		"clean_delete_fail",
		"clean_delete_success",
		"clean_unknown_command",
		"main_root_path_empty",
		"main_output_file_empty",
		"main_input_path_error",
		"main_input_path_not_dir",
		"main_open_scan_file_fail",
		"main_scan_info",
		"main_scan_time",
		"main_output_file_param_empty",
		"main_start_clean",
		"main_enter_params",
		"main_method_error",
		"main_delete_file_fail",
		"main_delete_scan_file_success",
	}

	// Test Chinese locale
	locales.LoadLocales("zh")
	for _, key := range keys {
		val := locales.GetMsg(key)
		if val == "" {
			t.Errorf("中文国际化key %s 返回空值", key)
		}
	}

	// Test English locale
	locales.LoadLocales("en")
	for _, key := range keys {
		val := locales.GetMsg(key)
		if val == "" {
			t.Errorf("英文国际化key %s 返回空值", key)
		}
	}
}
