package models

import (
	"errors"

	"tadmin/pkg/orm"
)

type Role struct {
	Id       int64  `json:"id" gorm:"primaryKey"` //角色ID
	Name     string `json:"name"`                 //角色名称
	Nickname string `json:"nickname"`             //显示名称
	Note     string `json:"note"`                 //备注
}

func (Role) TableName() string {
	return "sys_role"
}

func RoleGets(where string, args ...interface{}) ([]Role, error) {
	var objs []Role
	err := orm.DB.Where(where, args...).Find(&objs).Error
	if err != nil {
		return nil, errors.New("查询角色失败")
	}
	return objs, nil
}

// 增加角色
func (r *Role) Add() error {
	role, err := RoleGet("name = ?", r.Name)
	if err != nil {
		return errors.New("查询角色失败")
	}

	if role != nil {
		return errors.New("角色名称已存在")
	}

	return orm.DB.Create(r).Error
}

// 删除角色
func (r *Role) Del() error {
	return orm.DB.Delete(r).Error
}

// 更新角色
func (ug *Role) Update(selectField interface{}, selectFields ...interface{}) error {
	return orm.DB.Model(ug).Select(selectField, selectFields...).Updates(ug).Error
}

func RoleGet(where string, args ...interface{}) (*Role, error) {
	var lst []*Role
	err := orm.DB.Where(where, args...).Find(&lst).Error
	if err != nil {
		return nil, err
	}

	if len(lst) == 0 {
		return nil, nil
	}

	return lst[0], nil
}

func RoleCount(where string, args ...interface{}) (num int64, err error) {
	return Count(orm.DB.Model(&Role{}).Where(where, args...))
}

// RoleDel 删除角色
func RoleDel(id int64) error {

	target, err := RoleGet("id=?", id)

	if target.Name == "Admin" {
		return errors.New("Admin角色禁止更改")
	}

	var menu int64
	orm.DB.Model(&RoleOperation{}).Where("role_name = ?", target.Name).Count(&menu)
	var user int64
	orm.DB.Model(&User{}).Where("roles LIKE ?", target.Name).Count(&user)

	if menu > 0 && user > 0 {
		return errors.New("该角色存在关联的菜单和用户数据，删除失败")
	}

	err = target.Del()

	//err := orm.DB.Delete(&models.Role{}, "id = ?", id).Error

	//删除缓存
	//cache.Redis.DelByPattern(cache.CacheKeySysUserMenu)
	//cache.Redis.DelByPattern(cache.CacheKeySysUserPermission)

	return err
}
