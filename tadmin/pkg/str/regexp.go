package str

import (
	"regexp"
	"strings"
)

// regexp包实现了正则表达式匹配
// Compile解析并返回IP地址、邮箱正则表达式
var IPReg, _ = regexp.Compile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`)
var MailReg, _ = regexp.Compile(`\w[-._\w]*@\w[-._\w]*\.\w+`)

// IsMatch 检查字符串中是否存在匹配pattern的子序列
func IsMatch(s, pattern string) bool {
	match, err := regexp.Match(pattern, []byte(s))
	if err != nil {
		return false
	}

	return match
}

func IsIdentifier(s string, pattern ...string) bool {
	defpattern := "^[a-zA-Z0-9\\-\\_\\.]+$"
	if len(pattern) > 0 {
		defpattern = pattern[0]
	}

	return IsMatch(s, defpattern)
}

// IsMail 检查是否邮箱
func IsMail(s string) bool {
	return MailReg.MatchString(s)
}

// IsPhone 检查是否手机号码
func IsPhone(s string) bool {
	if strings.HasPrefix(s, "+") {
		return IsMatch(s[1:], `^\d{13}$`)
	} else {
		return IsMatch(s, `^\d{11}$`)
	}
}

// IsIP 检查是否IP地址
func IsIP(s string) bool {
	return IPReg.MatchString(s)
}

// Dangerous 检查是否包含常用字符
func Dangerous(s string) bool {
	if strings.Contains(s, "<") {
		return true
	}

	if strings.Contains(s, ">") {
		return true
	}

	if strings.Contains(s, "&") {
		return true
	}

	if strings.Contains(s, "'") {
		return true
	}

	if strings.Contains(s, "\"") {
		return true
	}

	if strings.Contains(s, "://") {
		return true
	}

	if strings.Contains(s, "../") {
		return true
	}

	return false
}
