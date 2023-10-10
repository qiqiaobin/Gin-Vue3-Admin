package models

type SysDict struct {
	Id       int64  `gorm:"primaryKey;comment:Id"`
	DictName string `gorm:"comment:字典名称"`
	DictCode string `gorm:"comment:字典代码"`
	DictType int    `gorm:"comment:字典类型（0int，1string）"`
	Status   int    `gorm:"comment:字典状态（0正常 1停用）"`
	Remark   string `json:"remark" gorm:"comment:备注"`
}

type SysDictItem struct {
	Id            int64  `gorm:"primaryKey;comment:选项id"`
	DictId        int64  `gorm:"comment:字典id"`
	DictItemName  string `gorm:"comment:选项名称"`
	DictItemValue string `gorm:"comment:选项值"`
	Status        int    `gorm:"comment:选项状态（0正常 1停用）"`
	Remark        string `json:"remark" gorm:"comment:备注"`
}
