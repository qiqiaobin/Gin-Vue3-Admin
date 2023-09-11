package system

import (
	"strings"
	"tadmin/models"
	"tadmin/pkg/ginx"

	"github.com/gin-gonic/gin"
)

type SysUserApi struct{}

// @Router /system/user/query [get]
func (sysUserApi *SysUserApi) Query(c *gin.Context) {
	limit := ginx.QueryInt(c, "limit", 20)
	query := ginx.QueryStr(c, "query", "")

	list, total, err := models.UserGets(query, limit, ginx.Offset(c, limit))

	ginx.Result(c, gin.H{
		"list":  list,
		"total": total,
		//"admin": user.IsAdmin(),
	}, err)
}

type userAddForm struct {
	Username string   `json:"username" binding:"required"`
	Password string   `json:"password" binding:"required"`
	Nickname string   `json:"nickname"`
	Phone    string   `json:"phone"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles" binding:"required"`
}

// @Router /system/user/add [post]
func (sysUserApi *SysUserApi) Add(c *gin.Context) {
	var addDto userAddForm
	c.ShouldBindJSON(&addDto)

	password, err := models.CryptoPass(addDto.Password)

	//一定要取到某个参数，取不到就panic
	//user := c.MustGet("user").(*models.User)
	u := models.User{
		Username: addDto.Username,
		Password: password,
		Nickname: addDto.Nickname,
		Phone:    addDto.Phone,
		Email:    addDto.Email,
		Roles:    strings.Join(addDto.Roles, " "),
		//CreateBy: user.Username,
		//UpdateBy: user.Username,
	}
	err = u.Add()

	ginx.Result(c, u, err)
}

type userProfileForm struct {
	Nickname string   `json:"nickname"`
	Phone    string   `json:"phone"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles"`
	//Contacts ormx.JSONObj `json:"contacts"`
}

// @Router /system/user/:id/update [put]
func (sysUserApi *SysUserApi) UserUpdate(c *gin.Context) {
	var f userProfileForm
	err := c.ShouldBindJSON(&f)

	//if len(f.Roles) == 0 {
	//	ginx.ResFailWithCode(c, 400, "用户角色为空")
	//}

	id := ginx.UrlParamInt64(c, "id")
	target, err := models.UserGetById(id)
	target.Nickname = f.Nickname
	target.Phone = f.Phone
	target.Email = f.Email
	target.Roles = strings.Join(f.Roles, " ")
	//target.Contacts = f.Contacts
	//target.UpdateBy = c.MustGet("username").(string)
	err = target.UpdateAllFields()

	ginx.Result(c, nil, err)
}

// @Router /system/user/:id [delete]
func (sysUserApi *SysUserApi) UserDel(c *gin.Context) {
	id := ginx.UrlParamInt64(c, "id")
	target, err := models.UserGetById(id)
	//ginx.Dangerous(err)

	//if target == nil {
	//	ginx.NewRender(c).Message(nil)
	//	return
	//}

	err = target.Del()

	ginx.Result(c, nil, err)
}

// @Router /system/user/:id [get]
func (sysUserApi *SysUserApi) GetDetail(c *gin.Context) {

	id := ginx.UrlParamInt64(c, "id")
	obj, err := models.UserGetById(id)
	//obj, err := userService.Detail(id)

	ginx.Result(c, obj, err)
}

// @Router /system/user/list [get]
func (sysUserApi *SysUserApi) List(c *gin.Context) {
	list, err := models.UserGetAll()
	ginx.Result(c, list, err)
}

type userPassword struct {
	Password string `json:"password" binding:"required"`
}

// @Router /system/user/:id/password [post]
func (sysUserApi *SysUserApi) PasswordRset(c *gin.Context) {
	var f userPassword
	err := c.ShouldBindJSON(&f)

	id := ginx.UrlParamInt64(c, "id")
	user, err := models.UserGetById(id)
	cryptoPass, err := models.CryptoPass(f.Password)
	user.UpdatePassword(cryptoPass)

	ginx.Result(c, nil, err)
}
