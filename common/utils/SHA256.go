package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// todo:SHA256 加密  哈希加密256位   转换为字符串后为64位
func Sha256Encrypt(str string) string {
	h := sha256.Sum256([]byte(str))
	return hex.EncodeToString(h[:])
}
