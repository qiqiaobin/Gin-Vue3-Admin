package system

import (
	"errors"
	"tadmin/models"
	"tadmin/pkg/ginx"
	"tadmin/pkg/logger"

	"github.com/gin-gonic/gin"
)

type SysRoleApi struct{}

// @Router /system/role/query [get]
func (*SysRoleApi) Query(c *gin.Context) {
	list, err := models.RoleGets("")
	total, err := models.RoleCount("")

	ginx.Result(c, gin.H{
		"list":  list,
		"total": total,
	}, err)
}

// @Router /system/role/add [post]
func (*SysRoleApi) Add(c *gin.Context) {
	var addData models.Role
	c.ShouldBindJSON(&addData)
	err := addData.Add()
	//err := roleService.Add(addDto)
	ginx.Result(c, nil, err)
}

// @Router /system/role/:id [put]
func (*SysRoleApi) Update(c *gin.Context) {
	var updateData models.Role
	var err error
	c.ShouldBindJSON(&updateData)

	oldRule, err := models.RoleGet("id=?", updateData.Id)
	if oldRule == nil {
		ginx.Result(c, nil, errors.New("角色不存在"))
	}

	if oldRule.Name != updateData.Name {
		// name changed, check duplication
		num, err := models.RoleCount("name=? and id<>?", updateData.Name, oldRule.Id)

		if num > 0 {
			ginx.Result(c, nil, err)
			logger.Errorf("角色已存在")
		}
	}

	oldRule.Name = updateData.Name
	oldRule.Note = updateData.Note

	err = oldRule.Update("name", "note")

	ginx.Result(c, nil, err)
}

// @Router /system/role/:id [delete]
func (*SysRoleApi) RoleDel(c *gin.Context) {
	var err error
	id := ginx.UrlParamInt64(c, "id")
	err = models.RoleDel(id)
	ginx.Result(c, nil, err)
}

// @Router /system/role/:id [get]
func (SysRoleApi) GetDetail(c *gin.Context) {

	id := ginx.UrlParamInt64(c, "id")
	obj, err := models.RoleGet("id=?", id)

	ginx.Result(c, obj, err)
}

// @Router /system/role/list [get]
func (*SysRoleApi) List(c *gin.Context) {
	list, err := models.RoleGets("")
	ginx.Result(c, list, err)
}

// @Router /system/role/:id/ops [put]
func (SysRoleApi) RoleBindOperation(c *gin.Context) {

	id := ginx.UrlParamInt64(c, "id")
	role, err := models.RoleGet("id=?", id)

	if role == nil {
		ginx.Result(c, nil, errors.New("角色不存在"))
	}

	var ops []int64
	c.ShouldBindJSON(&ops)

	err = models.RoleOperationBind(role.Name, ops)

	ginx.Result(c, role, err)
}
