package tools

import "strings"

// ConcatenatedStringBuilder 高效字符串拼接
// strings.Builder 是一个变长的字节缓存区，其内部使用slice来存储字节（buf []byte）
func ConcatenatedStringBuilder(str ...string) string {
	var builder strings.Builder
	for _, s := range str {
		builder.WriteString(s)
	}
	return builder.String()
}
