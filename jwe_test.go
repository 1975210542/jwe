package main

import (
	"fmt"
	"testing"
)

func Test_Jwe(t *testing.T) {
	header := NewHeader(ALG_RSA1_5, ENC_A128CBC_HS256)
	jsonHeader, err := utils.JsonEncode(header)
	fmt.Println("Header:", jsonHeader)

}
