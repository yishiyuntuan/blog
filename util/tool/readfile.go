package tool

import (
	"os"
)

// 读取文件返回字符串
func ReadFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return Byte2String(content)
}
