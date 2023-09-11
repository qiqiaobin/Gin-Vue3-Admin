package dto

// SysRoleQuery 查询
type SysRoleQuery struct {
	PageNum  int    `form:"pageNum"`                  // 页码
	PageSize int    `form:"pageSize"`                 // 每页数量
	Nickname string `json:"nickname" form:"nickname"` //角色名称
}

// SysRoleVo 输出对象
type SysRoleVo struct {
	Id       int64   `json:"id"`       //编号
	Nickname string  `json:"nickname"` //角色名称
	Name     string  `json:"name"`
	Note     string  `json:"note"`    //备注
	MenuIds  []int64 `json:"menuIds"` //角色菜单
}

type SysRoleAddDto struct {
	Nickname string  `json:"nickname"` //角色名称
	Name     string  `json:"name"`
	Note     string  `json:"note"`    //备注
	MenuIds  []int64 `json:"menuIds"` //角色菜单
}

// SysRoleUpdateDto 更新
type SysRoleUpdateDto struct {
	Id int64 `json:"id"` //编号
	SysRoleAddDto
}
