package scan

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/shootercheng/disk/internal/scan"
)

// TestScanFileByPath 测试扫描功能
func TestScanFileByPath(t *testing.T) {
	// 1. 创建临时根目录，测试结束后自动清理
	rootDir := t.TempDir()

	// 2. 构造测试文件系统结构
	// 结构如下:
	// rootDir/
	// ├── small.txt (10 bytes)
	// └── subdir/
	//     └── large.dat (100 bytes)

	subDir := filepath.Join(rootDir, "subdir")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}

	// 创建小文件
	smallFile := filepath.Join(rootDir, "small.txt")
	if err := os.WriteFile(smallFile, []byte("0123456789"), 0644); err != nil {
		t.Fatalf("Failed to create small file: %v", err)
	}

	// 创建大文件
	largeFile := filepath.Join(subDir, "large.dat")
	// 生成100字节内容
	content := make([]byte, 100)
	for i := range content {
		content[i] = byte(i % 256)
	}
	if err := os.WriteFile(largeFile, content, 0644); err != nil {
		t.Fatalf("Failed to create large file: %v", err)
	}

	// 3. 定义测试用例结构
	tests := []struct {
		name      string
		path      string
		threshold int64 // 阈值单位：字节
		fileSize  int64
	}{
		{
			name:      "scan root dir",
			path:      rootDir,
			threshold: 50,
			fileSize:  110,
		},
		{
			name:      "Non-existent path",
			path:      filepath.Join(rootDir, "non_existent"),
			threshold: 10,
			fileSize:  0,
		},
	}
	// 4. 执行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test for path: %s with threshold: %d", tt.path, tt.threshold)
			size := scan.ScanFileByPath(tt.path)
			if size != tt.fileSize {
				t.Fatal("scan file size error")
			}
		})
	}
}
