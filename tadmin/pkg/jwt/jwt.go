package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserAuthClaims struct {
	UserId   int64  //用户Id
	UserName string //用户账号
	NickName string //用户昵称
	jwt.RegisteredClaims
}

var secretKey = []byte("RcWJhezAYXjzVzZRmaX8B8uy2U3Hgg2p")

// GenerateToken 生成token
func GenerateToken(claim UserAuthClaims, expireTime time.Time) (string, error) {

	claim.RegisteredClaims = jwt.RegisteredClaims{
		// 签名时间
		IssuedAt: jwt.NewNumericDate(time.Now()),
		// 过期时间
		ExpiresAt: jwt.NewNumericDate(expireTime),
		// 签名颁发者
		Issuer: "gin-vue3-admin",
		// 签名主题
		Subject: "TF56",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString(secretKey)
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*UserAuthClaims, error) {
	// 解析 token string 拿到 token
	var claim UserAuthClaims
	token, err := jwt.ParseWithClaims(tokenStr, &claim, func(tk *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	claims, ok := token.Claims.(*UserAuthClaims)
	if !ok {
		return nil, err
	}

	return claims, err
}

// 创建用户token
func CreateToken(userId int64, username string) (string, error) {
	// 带权限创建令牌
	// 设置有效期，过期需要重新登录获取token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":      userId,
		"username": username,
		"exp":      time.Now().AddDate(0, 0, 1),
		//"exp":      time.Now().Add(time.Minute * time.Duration(conf.Config.JWT.ExpiresTime)).Unix(),
	})

	// 如果配置文件中的jwt key为空，则随机生成字符串
	//if conf.Config.JWT.SigningKey == "" {
	//	conf.Config.JWT.SigningKey = utils.GetRoundNumber(32)
	//	logger.Infof("config.yml未配置jwt.key, 随机生成key为: %s", conf.Config.JWT.SigningKey)
	//}
	// 使用自定义字符串加密 and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}

// 解析token，并返回登录者账号信息
func ParseToken2(tokenStr string) (*UserAuthClaims, error) {
	if tokenStr == "" {
		return nil, fmt.Errorf("未找到Token")
	}
	// Parse token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil || token == nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, ok := claims["uid"].(int64)
		if !ok {
			return nil, errors.New("failed to parse access_uuid from jwt")
		}

		return &UserAuthClaims{
			UserId:   uid,
			UserName: claims["username"].(string),
		}, nil
	}

	return nil, err
}
