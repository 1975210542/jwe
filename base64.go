package main

import (
	"encoding/base64"
	"strings"
)

func Base64Encode(json string) string {
	jh := base64.URLEncoding.EncodeToString(([]byte(json)))
	return strings.TrimRight(jh, "=")
}
