package service

import (
	"fmt"
	"strconv"

	"tadmin/models"
	"tadmin/models/dto"
	"tadmin/pkg/cache"
	"tadmin/pkg/orm"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

type SysAuthService struct{}

// GetUserId 获取当前用户Id
func GetUserId(c *gin.Context) int64 {
	claims, exists := c.Get("uid")

	uid, ok := claims.(int64)
	fmt.Println(uid)
	fmt.Println(exists)
	fmt.Println(ok)
	return uid
}

// GetUserInfo 获取用户信息
func (s *SysAuthService) GetUserInfo(c *gin.Context) (userInfo models.User, err error) {
	currentUserId := GetUserId(c)

	//查询用户信息
	err = orm.DB.Model(&models.User{}).Where("id = ?", currentUserId).Scan(&userInfo).Error
	if err != nil {
		return userInfo, err
	}

	//是否超级管理员
	//if jwt.IsAdmin(c) {
	if userInfo.IsAdmin() {
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
		orm.DB.Table("users").Select("roleMenu.menu_id").
			Joins("LEFT JOIN sys_role_menu roleMenu ON users.roles = roleMenu.role_name").
			Where("users.id = ?", userId).
			Group("roleMenu.menu_id").Find(&menuIds)

		orm.DB.Model(&models.SysMenu{}).
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

	var currentUserId = GetUserId(c)

	user, err := models.UserGetById(currentUserId)

	var cacheKey = cache.CacheKeySysUserMenu + strconv.FormatInt(currentUserId, 10)

	cache.Redis.Get(cacheKey, &authMenu)

	if authMenu == nil || len(authMenu) == 0 {
		var menuList []models.SysMenu
		//是否超级管理员
		if user.IsAdmin() {
			//超级管理员，拥有最高权限
			orm.DB.Model(&models.SysMenu{}).Where("menu_type IN (0,1) AND status = 0 ").Order("sort,id ASC").Scan(&menuList)
		} else {

			var menuIds []int64
			var userId = GetUserId(c)
			//用户拥有菜单id
			orm.DB.Table("users").Select("roleMenu.menu_id").
				Joins("LEFT JOIN sys_role_menu roleMenu ON users.roles = roleMenu.role_name").
				Where("users.id = ?", userId).
				Group("roleMenu.menu_id").Find(&menuIds)

			orm.DB.Model(&models.SysMenu{}).Where("id IN ? AND menu_type IN (0,1) AND status = 0 ",
				menuIds).Order("sort,id ASC").Scan(&menuList)
		}
		authMenu = BuildAuthMenuTree(menuList, 0)

		//设置缓存
		cache.Redis.Set(cacheKey, authMenu, cache.CacheTime)
	}

	return authMenu, err
}

// BuildAuthMenuTree 构建菜单树状
func BuildAuthMenuTree(menuList []models.SysMenu, parentId int64) (menus []dto.AuthMenuVo) {

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

// GetAuthMenu 根据角色id获取用户菜单路由Ids
func (s *SysAuthService) GetAuthMenuIds(roleId int64) ([]string, error) {

	var menuIds []string
	var err error

	var cacheKey = cache.CacheKeySysUserMenuIds + strconv.FormatInt(roleId, 10)

	cache.Redis.Get(cacheKey, &menuIds)

	role, err := models.RoleGet("id=?", roleId)

	if menuIds == nil || len(menuIds) == 0 {
		//是否超级管理员
		if role.Name == models.AdminRole {
			//超级管理员，拥有最高权限
			err = orm.DB.Table("sys_menu").Select("id").Where("menu_type IN (0,1) AND status = 0 ").Find(&menuIds).Error

		} else {
			//用户拥有菜单id
			err = orm.DB.Table("sys_role_menu").Select("menu_id").Where("role_name = ? ", role.Name).Find(&menuIds).Error

		}

		//设置缓存
		cache.Redis.Set(cacheKey, menuIds, cache.CacheTime)
	}

	fmt.Println(menuIds)
	return menuIds, err
}
