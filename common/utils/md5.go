package utils

import (
	"crypto/md5"
	"fmt"
)

// todo:加密32位
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}
