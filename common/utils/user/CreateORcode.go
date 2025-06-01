package user

import (
	"encoding/json"
	"fmt"
	"github.com/skip2/go-qrcode"
	"os"
	"time"
)

// MakeQrCode 生成二维码图片并保存到指定文件夹
func MakeQrCode(uid int, data interface{}) (string, error) {
	// 检查指定的文件夹是否存在，如果不存在则创建
	if _, err := os.Stat("../../srv/merchant_srv/images"); os.IsNotExist(err) {
		err = os.MkdirAll("../../srv/merchant_srv/images", 0755)
		if err != nil {
			return "", fmt.Errorf("创建文件夹 %s 时出错: %w", "../../srv/merchant_srv/images", err)
		}
	}

	// 获取当前时间并格式化
	now := time.Now().Format("20060102150405")
	// 生成文件名
	filename := fmt.Sprintf("user%d_%s.jpg", uid, now)
	// 拼接完整的文件路径
	filePath := fmt.Sprintf("%s/%s", "../../srv/merchant_srv/images", filename)

	// 将数据进行 JSON 序列化
	marshal, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("JSON 序列化数据时出错: %w", err)
	}
	fmt.Println(string(marshal))
	// 生成二维码并保存到指定文件

	return filename, qrcode.WriteFile(string(marshal), qrcode.Medium, 256, filePath)
}
