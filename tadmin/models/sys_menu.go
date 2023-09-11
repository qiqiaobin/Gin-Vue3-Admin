package models

import (
	"time"

	"gorm.io/gorm"
)

type SysMenu struct {
	Id            int64          `json:"id" gorm:"primaryKey;comment:编号"`
	ParentId      int64          `json:"parentId" gorm:"comment:父菜单ID"`
	MenuName      string         `json:"menuName" gorm:"comment:菜单名称"`
	MenuType      int            `json:"menuType" gorm:"comment:菜单类型（0菜单，1按钮）"`
	Path          string         `json:"path" gorm:"comment:路由地址"`
	Redirect      string         `json:"redirect" gorm:"comment:重定向路由地址"`
	Link          string         `json:"link" gorm:"comment:外部地址"`
	IsIframe      bool           `json:"isIframe" gorm:"comment:是否内嵌"`
	IsHide        bool           `json:"isHide" gorm:"comment:是否隐藏"`
	Component     string         `json:"component" gorm:"comment:组件路径"`
	Icon          string         `json:"icon" gorm:"comment:菜单图标"`
	Permission    string         `json:"permission" gorm:"comment:按钮权限"`
	RequestMethod string         `json:"requestMethod" gorm:"comment:请求方法"`
	Status        int            `json:"status" gorm:"comment:菜单状态（0正常 1停用）"`
	Sort          int            `json:"sort" gorm:"comment:排序"`
	Remark        string         `json:"remark" gorm:"comment:备注"`
	CreatedAt     time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt     time.Time      `json:"updatedAt" gorm:"comment:更新时间"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
