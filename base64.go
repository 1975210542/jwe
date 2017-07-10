package main

import (
	"encoding/base64"
	"strings"
)

//base64编码
func Base64Encode(json string) string {
	jh := base64.URLEncoding.EncodeToString(([]byte(json)))
	return strings.TrimRight(jh, "=")
	//	return jh
}

//base64解码
func Base64Decode(str string) ([]byte, error) {
	lenMod4 := len(str) % 4
	if lenMod4 > 0 {
		str = str + strings.Repeat("=", 4-lenMod4)
	}

	return base64.URLEncoding.DecodeString(str)
}
