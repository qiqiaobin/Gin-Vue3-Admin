package middleware

import (
	"log"
	"net/http"
	"runtime/debug"

	"tadmin/pkg/ginx"

	"github.com/gin-gonic/gin"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {

			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()

			ginx.FailWithCode(http.StatusInternalServerError, ErrorToString(r), c)
			c.Abort() //终止操作
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

func ErrorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
