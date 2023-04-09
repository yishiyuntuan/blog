package tool

import "unsafe"

// 字节切片转字符串
func Byte2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
