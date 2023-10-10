package dto

import (
	"time"
)

// SysDictQuery 查询
type SysDictQuery struct {
	PageNum  int    `form:"pageNum"`                 // 页码
	PageSize int    `form:"pageSize"`                // 每页数量
	DictName string `json:"dictName"form:"dictName"` //字典名称
}

// SysDictVo 输出对象
type SysDictVo struct {
	Id        int64        `json:"id"`             //编号
	DictName  string       `json:"dictName"`       //字典名称
	DictCode  string       `json:"dictCode"`       //字典代码
	DictType  int          `json:"dictType"`       //字典类型（0int，1string）
	Status    int          `json:"status"`         //字典状态（0正常 1停用）
	CreatedAt time.Time    `json:"createdAt"`      //创建时间
	UpdatedAt time.Time    `json:"updatedAt"`      //更新时间
	Remark    string       `json:"remark"`         //备注
	Items     []DictItemVo `json:"items" gorm:"-"` //选项
}

// SysDictAddDto 添加
type SysDictAddDto struct {
	DictName string `json:"dictName"` //字典名称
	DictCode string `json:"dictCode"` //字典代码
	DictType int    `json:"dictType"` //字典类型（0int，1string）
	Status   int    `json:"status"`   //字典状态（0正常 1停用）
	Remark   string `json:"remark"`   //备注
}

// SysDictUpdateDto 更新
type SysDictUpdateDto struct {
	Id int64 `json:"id"` //编号
	SysDictAddDto
}

// DictItemVo 输出对象
type DictItemVo struct {
	DictItemName  string      `json:"dictItemName"`  //选项名称
	DictItemValue interface{} `json:"dictItemValue"` //选项值
}

// SysDictItemVo 输出对象
type SysDictItemVo struct {
	Id            int64     `json:"id"`            //编号
	DictId        int64     `json:"dictId"`        //字典id
	DictItemName  string    `json:"dictItemName"`  //选项名称
	DictItemValue string    `json:"dictItemValue"` //选项值
	Status        int       `json:"status"`        //配置状态（0正常 1停用）
	CreatedAt     time.Time `json:"createdAt"`     //创建时间
	UpdatedAt     time.Time `json:"updatedAt"`     //更新时间
	Remark        string    `json:"remark"`        //备注
}
