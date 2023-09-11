package models

import (
	"tadmin/pkg/orm"
	"tadmin/pkg/slice"
)

type RoleOperation struct {
	Id        int64  `json:"id" gorm:"primaryKey;comment:角色和菜单关联id"`
	RoleName  string `json:"rolename"`  //角色名称
	MenuId    int64  `json:"menuId"`    //菜单id"`
	Operation string `json:"operation"` //操作路由
}

func (RoleOperation) TableName() string {
	return "sys_role_menu"
}

// RoleHasOperation 判断角色是否拥有权限路由
func RoleHasOperation(roles []string, operation string) (bool, error) {
	if len(roles) == 0 {
		return false, nil
	}

	var cnt int64
	err := orm.DB.Model(&RoleOperation{}).Where("operation = ? and role_name in ?", operation, roles).Count(&cnt).Error

	return cnt > 0, err
}

// OperationsOfRole  根据角色获取权限路由
func OperationsOfRole(roles []string) ([]string, error) {
	session := orm.DB.Model(&RoleOperation{}).Select("distinct(operation) as operation")

	if !slice.ContainsString(roles, AdminRole) {
		session = session.Where("role_name in ?", roles)
	}

	var ret []string
	err := session.Pluck("operation", &ret).Error
	return ret, err
}

// RoleOperationBind 角色权限路由绑定
func RoleOperationBind(roleName string, operation []int64) error {
	tx := orm.DB.Begin()

	if err := tx.Where("role_name = ?", roleName).Delete(&RoleOperation{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(operation) == 0 {
		return tx.Commit().Error
	}

	var ops []RoleOperation
	for _, op := range operation {
		ops = append(ops, RoleOperation{
			RoleName: roleName,
			MenuId:   op,
		})
	}

	if err := tx.Create(&ops).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
