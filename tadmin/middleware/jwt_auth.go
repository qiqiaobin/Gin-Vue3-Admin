package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"tadmin/pkg/ginx"
	"tadmin/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// JwtAuth 验证token
func JwtAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenStr := context.Request.Header.Get("Authorization")
		tokenStr = strings.Replace(tokenStr, "Bearer ", "", -1)

		if tokenStr == "" {
			ginx.FailWithCode(http.StatusUnauthorized, "无权限访问，请求未携带token", context)
			context.Abort() //结束后续操作
			return
		}

		// 解析token
		token, err := jwt.ParseToken(tokenStr)
		if err != nil {
			ginx.FailWithCode(http.StatusUnauthorized, err.Error(), context)
			context.Abort() //结束后续操作
			return
		}
		//是否过期

		// 获取 token 中的 claims
		claims, ok := token.Claims.(*jwt.UserAuthClaims)
		if !ok {
			context.JSON(http.StatusUnauthorized, "无权限访问，错误token")
			context.Abort()
			return
		}
		fmt.Println("获取 token 中的 claims")
		fmt.Println(claims)
		context.Set("claims", claims)
		context.Next()
	}
}
