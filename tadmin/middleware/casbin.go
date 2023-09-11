package middleware

import (
	"fmt"
	"strings"

	"tadmin/models"
	"tadmin/pkg/ginx"
	"tadmin/service"

	casbinUtil "github.com/casbin/casbin/v2/util"
	"github.com/gin-gonic/gin"
)

func Casbin() gin.HandlerFunc {
	return func(context *gin.Context) {
		//casbin需要定义model/policy
		//对于权限处理来说，此处主要目的：该用户的角色是否允许访问接口
		//最终决定使用KeyMatch2方法检查接口是否有权限访问即可

		//casbin毫无疑问是很强大的访问控制框架
		//但是个人觉得不应该为了用而用，先思考目的，将复杂问题简单化更重要
		//如有不同想法，欢迎交流 ✿(>‿◠)

		//是否超级管理
		uid := service.GetUserId(context)
		user, err := models.UserGetById(uid)
		if err != nil {
			return
		}
		if user.IsAdmin() {
			context.Next()
			return
		}
		permissionList := service.GetUserPermission(service.GetUserId(context))
		//是否允许访问
		var isAllow bool
		for i := 0; i < len(permissionList); i++ {
			var requestUrl = "/" + strings.Replace(permissionList[i].Permission, "_", "/", -1)
			if casbinUtil.KeyMatch2(strings.ToLower(context.Request.URL.Path), strings.ToLower(requestUrl)) &&
				context.Request.Method == permissionList[i].RequestMethod {
				isAllow = true
				break
			}
		}
		//检查接口是否有权限访问
		if !isAllow {
			fmt.Println("禁止访问")
			ginx.ResFailWithCode(context, 403, "禁止访问")
			context.Abort() //结束后续操作
			return
		}

		context.Next()
	}
}
