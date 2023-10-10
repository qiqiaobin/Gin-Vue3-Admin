package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"tadmin/pkg/orm"
	"tadmin/pkg/str"

	"gorm.io/gorm"
)

const (
	Dingtalk     = "dingtalk"
	Wecom        = "wecom"
	Feishu       = "feishu"
	FeishuCard   = "feishucard"
	Mm           = "mm"
	Telegram     = "telegram"
	Email        = "email"
	EmailSubject = "mailsubject"

	DingtalkKey = "dingtalk_robot_token"
	WecomKey    = "wecom_robot_token"
	FeishuKey   = "feishu_robot_token"
	MmKey       = "mm_webhook_url"
	TelegramKey = "telegram_robot_token"
)

// GORM 会直接填充更新时间、更新时间、删除时间
// 文档地址：https://gorm.io/zh_CN/docs/models.html#创建-x2F-更新时间追踪（纳秒、毫秒、秒、Time）
type User struct {
	Id       int64    `json:"id" gorm:"primaryKey"` //用户ID
	Username string   `json:"username"`             //用户账号
	Nickname string   `json:"nickname"`             //用户昵称
	Password string   `json:"-"`                    //密码
	Phone    string   `json:"phone"`                //手机号码
	Email    string   `json:"email"`                //用户邮箱
	Portrait string   `json:"portrait"`             //头像地址
	Roles    string   `json:"-"`                    // 这个字段写入数据库
	RolesLst []string `json:"roles" gorm:"-"`       // 这个字段和前端交互
	//Contacts orm.JSONObj `json:"contacts"`             // 内容为 map[string]string 结构
	//Maintainer int       `json:"maintainer"`           // 是否给管理员发消息 0:not send 1:send
	CreateAt   time.Time `json:"create_at"` //创建时间
	CreateBy   string    `json:"create_by"`
	UpdateAt   time.Time `json:"update_at"` //更新时间
	UpdateBy   string    `json:"update_by"`
	Admin      bool      `json:"admin" gorm:"-"`      // 方便前端使用
	Permission []string  `json:"permission" gorm:"-"` //权限
}

func (User) TableName() string {
	return "users"
}

/*
func (u *User) String() string {
	bs, err := u.Contacts.MarshalJSON()
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("<id:%d username:%s nickname:%s email:%s phone:%s contacts:%s>", u.Id, u.Username, u.Nickname, u.Email, u.Phone, string(bs))
}
*/

func (u *User) IsAdmin() bool {
	for i := 0; i < len(u.RolesLst); i++ {
		if u.RolesLst[i] == AdminRole {
			return true
		}
	}
	return false
}

// Verify 输入检查
func (u *User) Verify() error {
	//去掉用户名前后的空白字符
	u.Username = strings.TrimSpace(u.Username)

	if u.Username == "" {
		return errors.New("用户名为空")
	}

	if str.Dangerous(u.Username) {
		return errors.New("Username字符无效")
	}

	if str.Dangerous(u.Nickname) {
		return errors.New("Nickname字符无效")
	}

	if u.Phone != "" && !str.IsPhone(u.Phone) {
		return errors.New("手机号码无效")
	}

	if u.Email != "" && !str.IsMail(u.Email) {
		return errors.New("邮箱地址无效")
	}

	return nil
}

func (u *User) Add() error {
	user, err := UserGetByUsername(u.Username)
	if err != nil {
		return errors.New("查询用户失败")
	}

	if user != nil {
		return errors.New("此用户名已存在 请使用其他用户名")
	}

	now := time.Now()
	u.CreateAt = now
	u.UpdateAt = now
	return orm.DB.Create(u).Error
}

func (u *User) Update(selectField interface{}, selectFields ...interface{}) error {
	if err := u.Verify(); err != nil {
		return err
	}

	return orm.DB.Model(u).Select(selectField, selectFields...).Updates(u).Error
}

func (u *User) UpdateAllFields() error {
	if err := u.Verify(); err != nil {
		return err
	}

	u.UpdateAt = time.Now()
	return orm.DB.Model(u).Select("*").Updates(u).Error
}

func (u *User) UpdatePassword(password, updateBy string) error {
	return orm.DB.Model(u).Updates(map[string]interface{}{
		"password":  password,
		"update_at": time.Now(),
		"update_by": updateBy,
	}).Error
}

func (u *User) Del() error {
	return orm.DB.Transaction(func(tx *gorm.DB) error {
		//if err := tx.Where("user_id=?", u.Id).Delete(&UserGroupMember{}).Error; err != nil {
		//	return err
		//}

		if err := tx.Where("id=?", u.Id).Delete(&User{}).Error; err != nil {
			return err
		}

		return nil
	})
}

func (u *User) ChangePassword(oldpass, newpass string) error {
	_oldpass, err := CryptoPass(oldpass)
	if err != nil {
		return err
	}

	_newpass, err := CryptoPass(newpass)
	if err != nil {
		return err
	}

	if u.Password != _oldpass {
		return errors.New("Incorrect old password")
	}

	return u.UpdatePassword(_newpass, u.Username)
}

func PassLogin(username, pass string) (*User, error) {
	user, err := UserGetByUsername(username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("无效的用户名")
	}

	loginPass, err := CryptoPass(pass)
	if err != nil {
		return nil, err
	}

	if loginPass != user.Password {
		return nil, fmt.Errorf("登录密码错误")
	}

	return user, nil
}

func UserGets(query string, limit, offset int) ([]User, int64, error) {
	//session := orm.DB.Table("users").Limit(limit).Offset(offset).Order("id")
	session := orm.DB.Table("users").Limit(limit).Offset(offset).Order("id")
	//查询条件
	if query != "" {
		q := "%" + query + "%"
		session = session.Where("username like ? or nickname like ? or phone like ? or email like ?", q, q, q, q)
	}

	//总条数
	var total int64
	session.Count(&total)

	//查询数据
	var users []User
	err := session.Find(&users).Error
	if err != nil {
		return users, total, errors.New("查询用户失败")
	}

	for i := 0; i < len(users); i++ {
		users[i].RolesLst = strings.Fields(users[i].Roles)
		users[i].Admin = users[i].IsAdmin()
		users[i].Password = ""
	}

	return users, total, nil
}

func UserGet(where string, args ...interface{}) (*User, error) {
	var lst []*User
	err := orm.DB.Where(where, args...).Find(&lst).Error
	if err != nil {
		return nil, err
	}

	if len(lst) == 0 {
		return nil, nil
	}

	lst[0].RolesLst = strings.Fields(lst[0].Roles)
	lst[0].Admin = lst[0].IsAdmin()

	return lst[0], nil
}

func UserGetByUsername(username string) (*User, error) {
	return UserGet("username=?", username)
}

func UserGetById(id int64) (*User, error) {
	return UserGet("id=?", id)
}

func UserGetAll() ([]*User, error) {

	var lst []*User
	err := orm.DB.Find(&lst).Error
	if err == nil {
		for i := 0; i < len(lst); i++ {
			lst[i].RolesLst = strings.Fields(lst[i].Roles)
			lst[i].Admin = lst[i].IsAdmin()
		}
	}
	return lst, err
}

/*
func (u *User) ExtractToken(key string) (string, bool) {
	bs, err := u.Contacts.MarshalJSON()
	if err != nil {
		logger.Errorf("handle_notice: failed to marshal contacts: %v", err)
		return "", false
	}

	switch key {
	case Dingtalk:
		ret := gjson.GetBytes(bs, DingtalkKey)
		return ret.String(), ret.Exists()
	case Wecom:
		ret := gjson.GetBytes(bs, WecomKey)
		return ret.String(), ret.Exists()
	case Feishu, FeishuCard:
		ret := gjson.GetBytes(bs, FeishuKey)
		return ret.String(), ret.Exists()
	case Mm:
		ret := gjson.GetBytes(bs, MmKey)
		return ret.String(), ret.Exists()
	case Telegram:
		ret := gjson.GetBytes(bs, TelegramKey)
		return ret.String(), ret.Exists()
	case Email:
		return u.Email, u.Email != ""
	default:
		return "", false
	}
}
*/
