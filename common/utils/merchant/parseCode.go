package merchant

import (
	"encoding/json"
	"fmt"
	"github.com/tuotoo/qrcode"
	"os"
)

// ParseQrCode 解析二维码图片并将内容反序列化为指定类型
func ParseQrCode[T any](imageUrl string) (*T, error) {
	// 打开图片文件
	file, err := os.Open(fmt.Sprintf("images/%s", imageUrl))
	if err != nil {
		return nil, fmt.Errorf("无法打开文件: %w", err)
	}
	defer file.Close()

	// 解析二维码
	qrmatrix, err := qrcode.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("无法解析二维码: %w", err)
	}

	// 定义一个变量来存储反序列化的结果
	var result T
	// 将二维码内容反序列化为指定类型
	err = json.Unmarshal([]byte(qrmatrix.Content), &result)
	if err != nil {
		return nil, fmt.Errorf("无法反序列化二维码内容: %w", err)
	}

	return &result, nil
}
