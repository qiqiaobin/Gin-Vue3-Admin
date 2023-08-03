package service

import (
	"errors"
	"strconv"
	"time"

	"tadmin/model"
	"tadmin/model/dto"
	"tadmin/model/orm"
	"tadmin/pkg/cache"
	"tadmin/pkg/jwt"
	"tadmin/utils"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

type SysAuthService struct{}

// GenerateCaptcha 生成验证码
func (s *SysAuthService) GenerateCaptcha() (id string, b64s string, err error) {
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, store)
	return captcha.Generate()
}

// Login 用户登录
func (s *SysAuthService) Login(loginDto dto.LoginDto) (token string, err error) {

	if loginDto.CaptchaId == "" || loginDto.CaptchaCode == "" {
		return "", errors.New("验证码错误")
	}

	if !store.Verify(loginDto.CaptchaId, loginDto.CaptchaCode, true) {
		return "", errors.New("验证码错误")
	}

	var user model.SysUser

	err = orm.DB.Where("user_name = ?", loginDto.UserName).First(&user).Error
	if err != nil {
		return "", errors.New("账号不存在")
	}

	if user.Status == 1 {
		return "", errors.New("账号已被禁用")
	}

	pwd := utils.MD5(loginDto.Password + user.Salt)
	if pwd != user.Password {
		return "", errors.New("密码错误")
	}

	//生成token
	var claims = jwt.UserAuthClaims{
		UserId:   user.Id,
		UserName: user.UserName,
		NickName: user.NickName,
		UserType: user.UserType,
	}
	token, err = jwt.GenerateToken(claims, time.Now().AddDate(0, 0, 1))

	return token, err
}

// GetUserInfo 获取用户信息
func (s *SysAuthService) GetUserInfo(c *gin.Context) (userInfo dto.UserInfoVo, err error) {
	currentUserId := jwt.GetUserId(c)

	//查询用户信息
	err = orm.DB.Model(&model.SysUser{}).Where("id = ?", currentUserId).Scan(&userInfo).Error
	if err != nil {
		return userInfo, err
	}

	//是否超级管理员
	if jwt.IsSuperAdmin(c) {
		//所有权限
		userInfo.Permission = []string{"*_*_*"}
	} else {
		//用户权限
		userPermissions := GetUserPermission(currentUserId)
		for i := 0; i < len(userPermissions); i++ {
			userInfo.Permission = append(userInfo.Permission, userPermissions[i].Permission)
		}
	}
	return userInfo, err
}

// GetUserPermission 获取用户权限
func GetUserPermission(userId int64) (permissions []dto.UserPermissionVo) {

	var cacheKey = cache.CacheKeySysUserPermission + strconv.FormatInt(userId, 10)

	cache.Redis.Get(cacheKey, &permissions)
	if permissions == nil || len(permissions) == 0 {

		var menuIds []int64

		//用户拥有菜单id
		orm.DB.Table("sys_user_role userRole").Select("roleMenu.menu_id").
			Joins("LEFT JOIN sys_role_menu roleMenu ON userRole.role_id = roleMenu.role_id").
			Where("userRole.user_id = ?", userId).
			Group("roleMenu.menu_id").Find(&menuIds)

		orm.DB.Model(&model.SysMenu{}).
			Select("permission", "request_method").
			Where("id IN ? AND menu_type = 2 AND status = 0 ", menuIds).
			Scan(&permissions)

		//设置缓存
		cache.Redis.Set(cacheKey, permissions, cache.CacheTime)
	}
	return permissions
}

// GetAuthMenu 获取用户菜单路由
func (s *SysAuthService) GetAuthMenu(c *gin.Context) (authMenu []dto.AuthMenuVo, err error) {

	var currentUserId = jwt.GetUserId(c)

	var cacheKey = cache.CacheKeySysUserMenu + strconv.FormatInt(currentUserId, 10)

	cache.Redis.Get(cacheKey, &authMenu)

	if authMenu == nil || len(authMenu) == 0 {
		var menuList []model.SysMenu
		//是否超级管理员
		if jwt.IsSuperAdmin(c) {
			//超级管理员，拥有最高权限
			orm.DB.Model(&model.SysMenu{}).Where("menu_type IN (0,1) AND status = 0 ").Order("sort,id ASC").Scan(&menuList)
		} else {

			var menuIds []int64
			var userId = jwt.GetUserId(c)
			//用户拥有菜单id
			orm.DB.Table("sys_user_role userRole").Select("roleMenu.menu_id").
				Joins("LEFT JOIN sys_role_menu roleMenu ON userRole.role_id = roleMenu.role_id").
				Where("userRole.user_id = ?", userId).
				Group("roleMenu.menu_id").Find(&menuIds)

			orm.DB.Model(&model.SysMenu{}).Where("id IN ? AND menu_type IN (0,1) AND status = 0 ",
				menuIds).Order("sort,id ASC").Scan(&menuList)
		}
		authMenu = BuildAuthMenuTree(menuList, 0)

		//设置缓存
		cache.Redis.Set(cacheKey, authMenu, cache.CacheTime)
	}

	return authMenu, err
}

// BuildAuthMenuTree 构建菜单树状
func BuildAuthMenuTree(menuList []model.SysMenu, parentId int64) (menus []dto.AuthMenuVo) {

	var menuRouters []dto.AuthMenuVo
	for i := 0; i < len(menuList); i++ {
		var menu = menuList[i]
		if menu.ParentId == parentId {
			var menuRouter dto.AuthMenuVo
			menuRouter.Path = "/" + menu.Path
			menuRouter.Redirect = menu.Redirect
			menuRouter.Name = menu.Path
			menuRouter.Component = menu.Component
			menuRouter.Meta = dto.MenuRouterMetaVo{
				Title:    menu.MenuName,
				Icon:     menu.Icon,
				IsLink:   menu.Link,
				IsIframe: menu.IsIframe,
				IsHide:   menu.IsHide,
			}
			menuRouter.Children = BuildAuthMenuTree(menuList, menu.Id)
			menuRouters = append(menuRouters, menuRouter)
		}
	}
	return menuRouters
}

// UpdatePwd 修改密码
func (s *SysAuthService) UpdatePwd(c *gin.Context, updatePwdDto dto.UpdatePwdDto) error {
	var user model.SysUser
	currentUserId := jwt.GetUserId(c)
	err := orm.DB.Model(&model.SysUser{}).Where("id = ?", currentUserId).Scan(&user).Error
	if err != nil {
		return err
	}
	encryptPassword := utils.MD5(updatePwdDto.Password + user.Salt)
	if encryptPassword != user.Password {
		return errors.New("密码错误")
	}

	//密码盐
	salt := utils.GetRoundNumber(15)

	//加密 => MD5（密码+密码盐）
	pwd := utils.MD5(updatePwdDto.NewPassword + salt)

	//更新密码
	err = orm.DB.Model(&model.SysUser{}).Where("id = ?", user.Id).Updates(map[string]interface{}{
		"salt":     salt,
		"password": pwd,
	}).Error
	return err
}
