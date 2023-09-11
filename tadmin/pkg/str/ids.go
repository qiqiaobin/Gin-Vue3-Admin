package str

import (
	"fmt"
	"strconv"
	"strings"
)

// IdsInt64 将长字符串分割为int64数组，之间用sep来分割，默认用逗号分割
func IdsInt64(ids string, sep ...string) []int64 {
	if ids == "" {
		return []int64{}
	}

	s := ","
	if len(sep) > 0 {
		s = sep[0]
	}

	var arr []string

	if s == " " {
		//返回将字符串按照空白（unicode.IsSpace确定，可以是一到多个连续的空白字符）分割的多个字符串。如果字符串全部是空白或者是空字符串的话，会返回空切片。
		arr = strings.Fields(ids)
	} else {
		//用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。
		arr = strings.Split(ids, s)
	}

	count := len(arr)
	ret := make([]int64, 0, count)
	for i := 0; i < count; i++ {
		if arr[i] != "" {
			id, err := strconv.ParseInt(arr[i], 10, 64)
			if err == nil {
				ret = append(ret, id)
			}
		}
	}

	return ret
}

// IdsString 将int64数组转换为一个字符串，之间用sep来分隔，默认用逗号分隔
func IdsString(ids []int64, sep ...string) string {
	count := len(ids)
	arr := make([]string, count)
	for i := 0; i < count; i++ {
		arr[i] = fmt.Sprint(ids[i])
	}

	if len(sep) > 0 {
		return strings.Join(arr, sep[0])
	}

	return strings.Join(arr, ",")
}
