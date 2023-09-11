package ginx

import (
	"github.com/gin-gonic/gin"
)

// Response 响应信息结构体
type Response struct {
	Code int         `json:"code"` //状态码
	Data interface{} `json:"data"` //数据
	Msg  string      `json:"msg"`  //消息
}

// Result 响应函数
func Result(c *gin.Context, data interface{}, err error) {
	if err != nil {
		ResFail(c, err.Error())
	} else {
		ResSucc(c, data)
	}
}

// ResSucc 成功响应
func ResSucc(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code: 200,
		Msg:  "请求成功",
		Data: data,
	})
	return
}

// ResFail 失败响应
func ResFail(c *gin.Context, msg string) {
	c.JSON(200, Response{
		Code: 200,
		Msg:  msg,
		Data: nil,
	})
	return
}

// ResFail 失败响应
func ResFailWithCode(c *gin.Context, code int, msg string) {
	c.JSON(code, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
	return
}
