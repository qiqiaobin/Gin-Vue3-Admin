package system

import (
	"tadmin/models/dto"
	"tadmin/pkg/ginx"
	"tadmin/service"

	"github.com/gin-gonic/gin"
)

var menuService = &service.SysMenuService{}

type SysMenuApi struct{}

// Query
// @Tags 菜单
// @Summary 菜单查询
// @Security ApiKeyAuth
// @Param data query dto.SysMenuQuery true "请求参数"
// @Success 200 {object} response.JsonResult{data=response.PageResult{list=[]dto.SysMenuVo}}
// @Router /system/menu/query [get]
func (*SysMenuApi) Query(c *gin.Context) {
	var query dto.SysMenuQuery
	c.ShouldBindQuery(&query)

	list, total, err := menuService.Query(query)
	ginx.Result(c, ginx.PageResult{List: list, TotalCount: total}, err)
}

// Add
// @Tags 菜单
// @Summary 添加菜单
// @Security ApiKeyAuth
// @Param data body dto.SysMenuAddDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/menu/add [post]
func (*SysMenuApi) Add(c *gin.Context) {
	var addDto dto.SysMenuAddDto
	c.ShouldBindJSON(&addDto)

	err := menuService.Add(addDto)
	if err != nil {
		ginx.ResFail(c, err.Error())
	} else {
		ginx.ResSucc(c, nil)
	}
}

// Update
// @Tags 菜单
// @Summary 更新菜单
// @Security ApiKeyAuth
// @Param data body dto.SysMenuUpdateDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/menu/update [post]
func (*SysMenuApi) Update(c *gin.Context) {
	var updateDto dto.SysMenuUpdateDto
	c.ShouldBindJSON(&updateDto)

	err := menuService.Update(updateDto)
	if err != nil {
		ginx.ResFail(c, err.Error())
	} else {
		ginx.ResSucc(c, nil)
	}
}

// Delete
// @Tags 菜单
// @Summary 删除菜单
// @Security ApiKeyAuth
// @Param data body request.IdInfoDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /system/menu/delete [post]
func (*SysMenuApi) Delete(c *gin.Context) {
	var idInfoDto ginx.IdInfoDto
	c.ShouldBindJSON(&idInfoDto)

	err := menuService.Delete(idInfoDto.Id)
	if err != nil {
		ginx.ResFail(c, err.Error())
	} else {
		ginx.ResSucc(c, nil)
	}
}

// Detail
// @Tags 菜单
// @Summary 获取菜单详情
// @Security ApiKeyAuth
// @Param data query request.IdInfoDto true "菜单id"
// @Success 200 {object} response.JsonResult{data=dto.SysMenuVo}
// @Router /system/menu/detail [get]
func (*SysMenuApi) Detail(c *gin.Context) {
	var idInfoDto ginx.IdInfoDto
	c.ShouldBindQuery(&idInfoDto)

	obj, err := menuService.Detail(idInfoDto.Id)
	if err != nil {
		ginx.ResFail(c, err.Error())
	} else {
		ginx.ResSucc(c, obj)
	}
}

// List
// @Tags 菜单
// @Summary 菜单列表
// @Security ApiKeyAuth
// @Success 200 {object} response.JsonResult{data=[]dto.SysMenuVo}
// @Router /system/menu/list [get]
func (*SysMenuApi) List(c *gin.Context) {
	objs, err := menuService.List()
	if err != nil {
		ginx.ResFail(c, err.Error())
	} else {
		ginx.ResSucc(c, objs)
	}
}

// GetMenuTree
// @Tags 菜单
// @Summary 菜单树状
// @Security ApiKeyAuth
// @Success 200 {object} response.JsonResult{data=[]dto.SysMenuVo}
// @Router /system/menu/tree [get]
func (*SysMenuApi) GetMenuTree(c *gin.Context) {
	objs, err := menuService.GetMenuTree()
	if err != nil {
		ginx.ResFail(c, err.Error())
	} else {
		ginx.ResSucc(c, objs)
	}
}
