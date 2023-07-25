package model

import (
	"time"

	"gorm.io/gorm"
)

type SysRole struct {
	Id        int64          `json:"id" gorm:"primaryKey;comment:角色ID"`
	RoleName  string         `json:"roleName" gorm:"comment:角色名称"`
	RoleCode  string         `json:"roleCode" gorm:"comment:角色代码"`
	Status    int            `json:"status" gorm:"comment:菜单状态（0启用 1停用）"`
	Sort      int            `json:"sort" gorm:"comment:排序"`
	Remark    string         `json:"remark" gorm:"comment:备注"`
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (SysRole) TableName() string {
	return "sys_role"
}
