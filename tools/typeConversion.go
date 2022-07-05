package tools

import (
	"fmt"
	"strconv"

	"github.com/zeromicro/go-zero/core/jsonx"
)

// Float2Float float转float64并保留4位精度
func Float2Float(num float64, format string) float64 {
	floatNum, _ := strconv.ParseFloat(fmt.Sprintf(format, num), 64)
	return floatNum
}

// String2Int string 转 int
func String2Int(str string) int {
	intNum, _ := strconv.Atoi(str)
	return intNum
}

// String2Int64 string 转 int64
func String2Int64(str string) int64 {
	intNum, _ := strconv.Atoi(str)
	return int64(intNum)
}

// String2Float64 string 转 float64
func String2Float64(str string) float64 {
	floatNum, _ := strconv.ParseFloat(str, 64)
	return floatNum
}

// String2Bool 字符串转bool
func String2Bool(str string) bool {
	boolVal, _ := strconv.ParseBool(str)
	return boolVal
}

// Decimal float64 保留 两位小数
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

// JsonTostring
func JsonTostring(v interface{}) string {
	bytes, err := jsonx.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// SearchParam 解析查询参数
func SearchParam(s string) (str []string, value []interface{}) {
	var res = map[string]interface{}{}
	jsonx.UnmarshalFromString(s, &res)
	for k, v := range res {
		if k != "page" && k != "limit" && k != "order" && k != "sort" {
			str = append(str, k+"=")
			value = append(value, v)
		}
	}
	return
}

// string to v
func StringToInterface(s string, v interface{}) error {
	return jsonx.Unmarshal([]byte(s), v)
}
