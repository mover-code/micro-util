package tools

import (
	"fmt"
	"strconv"
)

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
func StringToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}
func StringToFloat64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
func Float64ToString(f float64) string {
	return fmt.Sprintf("%0.2f", f)
}
