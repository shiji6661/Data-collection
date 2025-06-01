package user

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// 定义 Base62 编码所需的字符集
const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Base62Encode 将整数转换为 Base62 编码的字符串
func Base62Encode(num uint64) string {
	if num == 0 {
		return string(base62Chars[0])
	}
	encoded := ""
	for num > 0 {
		remainder := num % 62
		encoded = string(base62Chars[remainder]) + encoded
		num /= 62
	}
	return encoded
}

// Base62Decode 将 Base62 编码的字符串解码为整数
func Base62Decode(encoded string) (uint64, error) {
	var decoded uint64
	for _, char := range encoded {
		pos := -1
		for i, c := range base62Chars {
			if c == char {
				pos = i
				break
			}
		}
		if pos == -1 {
			return 0, fmt.Errorf("invalid character in base62 string: %c", char)
		}
		decoded = decoded*62 + uint64(pos)
	}
	return decoded, nil
}

// GenerateRandomString 生成指定长度的随机字符串
func GenerateRandomString(length int) (string, error) {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(base62Chars))))
		if err != nil {
			return "", err
		}
		result[i] = base62Chars[n.Int64()]
	}
	return string(result), nil
}

// GenerateInviteCode  todo 根据用户 ID 生成邀请码
func GenerateInviteCode(userID uint64) (string, error) {
	encodedID := Base62Encode(userID)
	randomSuffix, err := GenerateRandomString(4)
	if err != nil {
		return "", err
	}
	return encodedID + randomSuffix, nil
}

// ParseInviteCode  todo  从邀请码中解析出用户 ID
func ParseInviteCode(inviteCode string) (uint64, error) {
	if len(inviteCode) < 4 {
		return 0, fmt.Errorf("invalid invite code: too short")
	}
	encodedID := inviteCode[:len(inviteCode)-4]
	return Base62Decode(encodedID)
}
