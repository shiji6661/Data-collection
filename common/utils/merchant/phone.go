package merchant

import (
	"errors"
	"fmt"
	"regexp"
)

func Phone(MerchantPhone string) error {
	pattern := `^1[3-9]\d{9}$`
	match, _ := regexp.MatchString(pattern, MerchantPhone)
	if match {
		fmt.Println("手机号码格式正确")
	} else {
		return errors.New("手机号码格式错误")
	}
	return nil
}
