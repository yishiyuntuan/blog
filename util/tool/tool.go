package tool

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/kataras/iris/v12"
)

// SplitToIntList 字符串切割成int数组
func SplitToIntList(str string, sep string) (intList []int) {
	if str == "" {
		return
	}
	strList := strings.Split(str, sep)
	if len(strList) == 0 {
		return
	}
	for _, item := range strList {
		if item == "" {
			continue
		}
		val, err := strconv.ParseInt(item, 10, 32)
		if err != nil {
			// logs.CtxError(ctx, "ParseInt fail, err=%v, str=%v, sep=%v", err, str, sep) // 此处打印出log报错信息
			continue
		}
		intList = append(intList, int(val))
	}
	return
}

/*
PageTool 分页通用获取
参数1 pageSize 分页最大数
参数2 pageNum 分页偏移量
返回值
*/
func PageTool(ctx iris.Context) (int, int) {
	pageSize, err := ctx.URLParamInt("pageSize")
	if err != nil {
		pageSize = 0
	}
	pageNum, err := ctx.URLParamInt("pageNum")
	if err != nil {
		pageNum = 0
	}
	if pageSize <= 0 || pageSize > 20 {
		pageSize = 10
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	return pageSize, pageNum
}

func PageIds(pageNum, pageSize int, dataIds []string) (ids []string) {
	n := len(dataIds)
	if pageNum > n || (pageNum-1)*pageSize > n {
		ids = nil
	} else if pageNum == -1 {
		ids = dataIds
	} else if pageNum*pageSize > n {
		ids = dataIds[(pageNum-1)*pageSize:]
	} else {
		ids = dataIds[(pageNum-1)*pageSize : pageNum*pageSize]
	}
	return
}

// VerifyFormat 校验格式
func VerifyFormat(format string, data string) bool {
	if data == "" {
		return false
	}
	reg := regexp.MustCompile(format)
	return reg.MatchString(data)
}
