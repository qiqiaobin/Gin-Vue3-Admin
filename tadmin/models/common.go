package models

import (
	"tadmin/pkg/str"

	"gorm.io/gorm"
)

const AdminRole = "Admin"

func Count(tx *gorm.DB) (int64, error) {
	var cnt int64
	err := tx.Count(&cnt).Error
	return cnt, err
}

// CryptoPass 密码加密
func CryptoPass(raw string) (string, error) {
	salt := "7f07b0704579ed306a62ee0aef469a94"

	return str.MD5(salt + "<-*Uk30^96eY*->" + raw), nil
}
