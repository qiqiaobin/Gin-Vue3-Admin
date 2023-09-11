package middleware

import (
	"fmt"
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
			ginx.ResFailWithCode(context, 401, "无权限访问，请求未携带token")
			context.Abort() //结束后续操作
			return
		}

		// 解析token
		claims, err := jwt.ParseToken(tokenStr)
		if err != nil {
			context.JSON(401, "无权限访问，错误token")
			context.Abort()
			return
		}
		fmt.Println("获取 token 中的 claims")
		fmt.Println(claims)
		context.Set("uid", claims.UserId)
		context.Set("username", claims.UserName)
		context.Next()
	}
}
