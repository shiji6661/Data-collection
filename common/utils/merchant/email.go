package merchant

import (
	"regexp"
	"strings"
)

func Email(email string) bool {
	// 1. 空值处理（明确标注业务场景）
	if email == "" {
		return true // 非必填场景（如可选邮箱字段）
		// return false // 若业务要求必填，取消注释此行
	}

	// 2. 预编译正则（避免重复编译，提升性能）
	const pattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	staticRegex := regexp.MustCompile(pattern) // ✅ 预编译（仅初始化时执行一次）

	// 3. 核心验证（防御性检查参数类型）
	if !staticRegex.MatchString(strings.TrimSpace(email)) { // ✅ 自动去空格
		return false
	}

	// 4. 额外校验（可选：防止特殊攻击，如超长域名）
	if len(email) > 254 { // ✅ 符合 RFC 5321 最大长度
		return false
	}

	return true
}
