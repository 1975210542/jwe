package main

import (
	//	"fmt"

	"strings"
)

type JwtToken struct {
	Header map[string]interface{}
	Claims Claims
	Method EncryptionMethod
}

func NewWithClaims(method EncryptionMethod) *JwtToken {
	return newWithClaims(method, MapClaims{"username": "gaoqiankun"})
}

func newWithClaims(method EncryptionMethod, claims Claims) *JwtToken {
	return &JwtToken{
		Header: map[string]interface{}{
			"type": "JWT",
			"alg":  method.GetName(),
		},
		Claims: claims,
		Method: method,
	}
}

//将用"."连接的header和payload进行数字签名
func (j *JwtToken) SignedToken(key interface{}) (string, error) {
	var err error
	var str string
	var stoken []byte
	if str, err = j.SignedString(); err != nil {
		return "", err
	}
	if stoken, err = j.Method.Encrypt([]byte(str), key); err != nil {
		return "", err
	}
	return strings.Join([]string{str, string(stoken)}, "."), nil
}

//将header和payload进行Base64编码，并用".",连接
func (j *JwtToken) SignedString() (string, error) {
	var err error
	part := make([]string, 2)
	for i, _ := range part {
		var jsonValue string
		if i == 0 {
			if jsonValue, err = JsonEncode(j.Header); err != nil {
				return "", err
			}
		} else {
			if jsonValue, err = JsonEncode(j.Claims); err != nil {
				return "", err
			}
		}
		part[i] = Base64Encode(jsonValue)
	}

	return strings.Join(part, "."), nil
}
