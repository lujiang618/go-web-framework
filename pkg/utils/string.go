package utils

import "strings"

// 字符串拼接
func StringSplice(strList ...string) string {
	if len(strList) == 0 {
		return ""
	}

	var builder strings.Builder
	for _, str := range strList {
		builder.WriteString(str)
	}

	return builder.String()
}
