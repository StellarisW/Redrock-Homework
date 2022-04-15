package utils

import (
	"strconv"
	"strings"
)

func UnicodeMarshal(rawString string) string {
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(rawString), `\\u`, `\u`, -1))
	return str
}
