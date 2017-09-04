package utils

import "strings"

func Trim(s string) string {
	return strings.Trim(s, "\n ")
}
