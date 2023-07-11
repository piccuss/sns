package utils

import (
	"fmt"
	"strconv"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

// 以指定的charset将byte[]转换成string
func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}

// 将string转换成int
func ParseInt(num string) int {
	t, _ := strconv.Atoi(num)
	return t
}

// 将string转换成float64
func ParseFloat64(num string) float64 {
	t, _ := strconv.ParseFloat(num, 64)
	return t
}

// 将float64转换成固定三位的小数
func ConvertToFixDecimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", value), 64)
	return value
}
