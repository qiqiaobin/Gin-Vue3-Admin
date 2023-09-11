package ginx

import (
	"strconv"
	"tadmin/pkg/logger"

	"github.com/gin-gonic/gin"
)

// UrlParamStr 从URL路径中获取参数，输出字符串
func UrlParamStr(c *gin.Context, field string) string {
	val := c.Param(field)

	if val == "" {
		logger.Errorf("url参数[%s]为空", field)
	}

	return val
}

// UrlParamInt64 从URL路径中获取参数，输出Int64
func UrlParamInt64(c *gin.Context, field string) int64 {
	strval := UrlParamStr(c, field)
	intval, err := strconv.ParseInt(strval, 10, 64)
	if err != nil {
		logger.Errorf("不能转换 %s 为Int64类型", strval)
	}

	return intval
}

func UrlParamInt(c *gin.Context, field string) int {
	strval := UrlParamStr(c, field)
	intval, err := strconv.Atoi(strval)
	if err != nil {
		logger.Errorf("不能转换 %s 为Int类型", strval)
	}

	return intval
}

// QueryStr获取查询参数，不存在则返回默认值
func QueryStr(c *gin.Context, key string, defaultVal ...string) string {
	val := c.Query(key)
	if val != "" {
		return val
	}

	if len(defaultVal) == 0 {
		logger.Errorf("查询参数[%s]不能为空", key)
	}

	return defaultVal[0]
}

// QueryInt获取查询参数中指定参数值，并转为int
func QueryInt(c *gin.Context, key string, defaultVal ...int) int {
	strv := c.Query(key)
	if strv != "" {
		intv, err := strconv.Atoi(strv)
		if err != nil {
			logger.Errorf("不能转换[%s]为Int类型", strv)
		}
		return intv
	}

	if len(defaultVal) == 0 {
		logger.Errorf("查询参数[%s]不能为空", key)
	}

	return defaultVal[0]
}

// QueryInt64获取查询参数中指定参数值，并转为int64
func QueryInt64(c *gin.Context, key string, defaultVal ...int64) int64 {
	strv := c.Query(key)
	if strv != "" {
		intv, err := strconv.ParseInt(strv, 10, 64)
		if err != nil {
			logger.Errorf("不能转换[%s]为Int64类型", strv)
		}
		return intv
	}

	if len(defaultVal) == 0 {
		logger.Errorf("查询参数[%s]不能为空", key)
	}

	return defaultVal[0]
}
