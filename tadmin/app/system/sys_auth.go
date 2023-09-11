package system

import (
	model "tadmin/models"
	"tadmin/pkg/ginx"
	"tadmin/pkg/jwt"
	"tadmin/pkg/logger"
	"tadmin/service"
	"time"

	"github.com/gin-gonic/gin"
	captcha "github.com/mojocn/base64Captcha"
)

var store = captcha.DefaultMemStore

var authService = &service.SysAuthService{}

type SysAuthApi struct{}

// 生成图形验证码
func (*SysAuthApi) GenerateCaptcha(c *gin.Context) {
	var driver = captcha.NewDriverMath(80, 240, 0, captcha.OptionShowSlimeLine, nil, nil, []string{"wqy-microhei.ttc"})
	cc := captcha.NewCaptcha(driver, store)
	//data:image/png;base64
	id, b64s, err := cc.Generate()

	ginx.Result(c, gin.H{
		"imgdata":   b64s,
		"captchaid": id,
	}, err)
}

type LoginDto struct {
	UserName    string `json:"userName"`    //用户账号
	Password    string `json:"password"`    // 密码
	CaptchaCode string `json:"captchaCode"` // 验证码
	CaptchaId   string `json:"captchaId"`   // 验证码ID
}

// @Router /system/auth/login [post]
func (*SysAuthApi) Login(c *gin.Context) {

	var loginDto LoginDto
	c.ShouldBindJSON(&loginDto)
	//token, err := authService.Login(loginDto)

	if loginDto.CaptchaId == "" || loginDto.CaptchaCode == "" {
		logger.Error("验证码错误")
		return
	}

	if !store.Verify(loginDto.CaptchaId, loginDto.CaptchaCode, true) {
		logger.Error("验证码错误")
		return
	}

	var user model.User

	user, err := authService.PassLogin(loginDto.UserName, loginDto.Password)
	//生成token
	var claims = jwt.UserAuthClaims{
		UserId:   user.Id,
		UserName: user.Username,
		NickName: user.Nickname,
	}
	token, err := jwt.GenerateToken(claims, time.Now().AddDate(0, 0, 1))

	//token, err := jwt.CreateToken(user.Id, user.Username)

	ginx.Result(c, token, err)
}

// GetUserInfo
// @Tags 授权
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @Success 200 {object} response.JsonResult{data=models.User}
// @Router /system/auth/userInfo [get]
func (*SysAuthApi) GetUserInfo(c *gin.Context) {
	userInfo, err := authService.GetUserInfo(c)
	ginx.Result(c, userInfo, err)

}

// GetAuthMenu
// @Tags 授权
// @Summary 获取用户菜单（树状）
// @Security ApiKeyAuth
// @Success 200 {object} response.JsonResult{data=[]dto.AuthMenuVo}
// @Router /system/auth/menu [get]
func (*SysAuthApi) GetAuthMenu(c *gin.Context) {
	authMenu, err := authService.GetAuthMenu(c)
	ginx.Result(c, authMenu, err)
}

// @Router /system/auth/:rolename/menuids [get]
func (*SysAuthApi) GetAuthMenuIds(c *gin.Context) {
	//rolename := ginx.QueryStr(c, "rolename", "")
	//id := ginx.UrlParamInt64(c, "roleId")
	var idInfoDto ginx.IdInfoDto
	c.ShouldBindQuery(&idInfoDto)
	authMenuids, err := authService.GetAuthMenuIds(idInfoDto.Id)
	ginx.Result(c, authMenuids, err)
}
