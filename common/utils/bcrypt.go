package utils

import "golang.org/x/crypto/bcrypt"

// todo:大约为60位
// todo:自带加盐
//
//todo:bcrypt 的输出长度并不固定，它包含了盐值、算法版本和哈希值本身等信息
func BcryptEncrypt(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(bytes), err
}
