package pkg

import (
	"strings"
)

// ConvertToChineseCurrency 将数字金额转换为中文大写
func ConvertToChineseCurrency(amount float64) string {
	if amount < 0 || amount > 9999999999999.99 {
		return "金额超出范围"
	}

	integerPart := int64(amount)
	decimalPart := int64((amount - float64(integerPart)) * 100)

	chineseDigits := []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}
	chineseUnits := []string{"", "拾", "佰", "仟", "万", "拾", "佰", "仟", "亿", "拾", "佰", "仟", "万"}

	var result strings.Builder

	// 处理整数部分
	if integerPart == 0 {
		result.WriteString(chineseDigits[0])
	} else {
		index := 0
		for integerPart > 0 {
			digit := integerPart % 10
			if digit != 0 {
				result.WriteString(chineseUnits[index])
				result.WriteString(chineseDigits[digit])
			} else {
				// 处理零的情况
				if index == 4 || index == 8 {
					result.WriteString(chineseUnits[index])
				} else if integerPart/10%10 != 0 {
					result.WriteString(chineseDigits[digit])
				}
			}
			integerPart /= 10
			index++
		}
		// 反转字符串
		runes := []rune(result.String())
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		result.Reset()
		result.WriteString(string(runes))
		result.WriteString("元")
	}

	// 处理小数部分
	if decimalPart == 0 {
		result.WriteString("整")
	} else {
		jiao := decimalPart / 10
		fen := decimalPart % 10

		if jiao != 0 {
			result.WriteString(chineseDigits[jiao])
			result.WriteString("角")
		} else if fen != 0 {
			result.WriteString(chineseDigits[0])
		}

		if fen != 0 {
			result.WriteString(chineseDigits[fen])
			result.WriteString("分")
		}
	}

	return result.String()
}
