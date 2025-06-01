package merchant

import (
	"regexp"
)

// 识别身份类型
func IdentityType(username string) string {
	if regexp.MustCompile(`^1[3-9]\d{9}$`).MatchString(username) { // 手机号
		return "phone"
	} else if regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`).MatchString(username) { // 邮箱
		return "email"
	} else if len(username) >= 1 { // 简单用户名验证
		return "username"
	}
	return ""
}
