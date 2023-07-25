package dto

import (
	"time"
)

// SysRoleQuery 查询
type SysRoleQuery struct {
	PageNum  int    `form:"pageNum"`                  // 页码
	PageSize int    `form:"pageSize"`                 // 每页数量
	RoleName string `json:"roleName" form:"roleName"` //角色名称
}

// SysRoleVo 输出对象
type SysRoleVo struct {
	Id        int64     `json:"id"`        //编号
	RoleName  string    `json:"roleName"`  //角色名称
	RoleCode  string    `json:"roleCode"`  //角色代码
	Status    int       `json:"status"`    //菜单状态（0正常 1停用）
	Sort      int       `json:"sort"`      //排序
	Remark    string    `json:"remark"`    //备注
	CreatedAt time.Time `json:"createdAt"` //创建时间
	UpdatedAt time.Time `json:"updatedAt"` //更新时间
	MenuIds   []int64   `json:"menuIds"`   //角色菜单
}

type SysRoleAddDto struct {
	RoleName string  `json:"roleName"` //角色名称
	RoleCode string  `json:"roleCode"` //角色代码
	Status   int     `json:"status"`   //菜单状态（0正常 1停用）
	Sort     int     `json:"sort"`     //排序
	Remark   string  `json:"remark"`   //备注
	MenuIds  []int64 `json:"menuIds"`  //角色菜单
}

// SysRoleUpdateDto 更新
type SysRoleUpdateDto struct {
	Id int64 `json:"id"` //编号
	SysRoleAddDto
}
