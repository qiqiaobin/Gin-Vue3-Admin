package sys

import (
	"tadmin/app/system"

	"github.com/gin-gonic/gin"
)

// SystemRouter 系统路由
// 建议：
// 当接口数量超出一定范围，可适当分类或将复杂模块独立文件整理
// 路由规范：业务/模块/操作，比如：mall/user/add
type SystemRouter struct{}

var (
	authApi = &system.SysAuthApi{}
	menuApi = &system.SysMenuApi{}
	userApi = &system.SysUserApi{}
	roleApi = &system.SysRoleApi{}
)

// InitPublicRouter 初始化公开路由
func (s *SystemRouter) InitPublicRouter(routerGroup *gin.RouterGroup) {

	//权限
	authRouter := routerGroup.Group("/system/auth")
	{
		authRouter.GET("captcha", authApi.GenerateCaptcha) // 生成验证码
		authRouter.POST("login", authApi.Login)            // 用户登录
	}
}

// InitPrivateRouter 初始化私有路由
func (s *SystemRouter) InitPrivateRouter(routerGroup *gin.RouterGroup) {

	//权限
	authRouter := routerGroup.Group("/system/auth")
	{
		authRouter.GET("userInfo", authApi.GetUserInfo)    // 获取用户信息
		authRouter.GET("menu", authApi.GetAuthMenu)        // 获取用户菜单（树状）
		authRouter.GET("/menuids", authApi.GetAuthMenuIds) // 获取用户菜单ids
	}

	//用户
	userRouter := routerGroup.Group("/system/user")
	{
		userRouter.GET("query", userApi.Query)               // 用户分页查询
		userRouter.POST("add", userApi.Add)                  // 添加用户
		userRouter.PUT(":id/update", userApi.UserUpdate)     // 更新用户
		userRouter.DELETE(":id", userApi.UserDel)            // 删除用户
		userRouter.GET("list", userApi.List)                 // 用户列表
		userRouter.GET(":id", userApi.GetDetail)             // 用户详情
		userRouter.PUT(":id/password", userApi.PasswordRset) // 重置密码
	}

	//菜单
	menuRouter := routerGroup.Group("/system/menu")
	{
		menuRouter.GET("query", menuApi.Query)      // 菜单分页查询
		menuRouter.POST("add", menuApi.Add)         // 添加菜单
		menuRouter.POST("update", menuApi.Update)   // 更新菜单
		menuRouter.POST("delete", menuApi.Delete)   // 删除菜单
		menuRouter.GET("detail", menuApi.Detail)    // 菜单详情
		menuRouter.GET("list", menuApi.List)        // 菜单列表
		menuRouter.GET("tree", menuApi.GetMenuTree) // 菜单树状
	}

	//角色
	roleRouter := routerGroup.Group("/system/role")
	{
		roleRouter.GET("query", roleApi.Query)                  // 角色查询
		roleRouter.POST("add", roleApi.Add)                     // 添加角色
		roleRouter.POST("update", roleApi.Update)               // 更新角色
		roleRouter.PUT(":id/update", roleApi.RoleBindOperation) // 更新用户
		roleRouter.DELETE(":id", roleApi.RoleDel)               // 删除角色
		roleRouter.GET(":id", roleApi.GetDetail)                // 用户详情
		roleRouter.GET("list", roleApi.List)                    // 角色列表
	}
}
