package user

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// 字符集，用于生成邀请码中的随机部分
const charset = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"

// 分隔符，用于分隔用户 ID 和随机部分
const separator = "-"

// GenerateInviteCodeWithSeparator 生成包含用户 ID 的邀请码
// userId 是用户的唯一标识
// 返回值是生成的邀请码，如果生成过程中出现错误则返回空字符串
func GenerateInviteCodeWithSeparator(userId int64) string {
	// 将用户 ID 转换为字符串
	userIdStr := strconv.FormatInt(userId, 10)
	// 计算邀请码中随机部分的长度
	remainingLength := 8 - len(userIdStr) - len(separator)
	// 如果剩余长度小于 0，说明用户 ID 过长，截取前 8 位作为邀请码
	if remainingLength < 0 {
		return userIdStr[:8]
	}
	// 初始化一个字节切片，用于存储随机部分
	inviteCode := make([]byte, remainingLength)
	// 循环生成随机部分
	for i := range inviteCode {
		// 从字符集中随机选择一个字符
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return ""
		}
		// 将随机字符添加到邀请码中
		inviteCode[i] = charset[n.Int64()]
	}
	// 拼接用户 ID、分隔符和随机部分，返回完整的邀请码
	return userIdStr + separator + string(inviteCode)
}

// ExtractUserIdFromCodeWithSeparator 从邀请码中提取用户 ID
// inviteCode 是生成的邀请码
// 返回值是提取的用户 ID 和可能出现的错误
func ExtractUserIdFromCodeWithSeparator(inviteCode string) (int64, error) {
	// 使用分隔符分割邀请码
	parts := strings.Split(inviteCode, separator)
	// 如果分割后得到两部分，尝试将第一部分转换为用户 ID
	if len(parts) == 2 {
		userId, err := strconv.ParseInt(parts[0], 10, 64)
		if err == nil {
			return userId, nil
		}
	}
	// 如果无法提取用户 ID，返回错误信息
	return 0, fmt.Errorf("无法从邀请码中提取用户 ID")
}
