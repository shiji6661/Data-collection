package user

import "regexp"

// todo 正则表达式验证中国手机号格式
func ValidateMobile(mobile string) bool {
	pattern := `^1[3-9]\d{9}$`
	match, _ := regexp.MatchString(pattern, mobile)
	return match
}
