package main

import (
	"fmt"
	_ "jwe/aes"
	_ "jwe/hma"
	_ "jwe/rsa"
	"testing"
)

func Test_Jwe(t *testing.T) {
	args := make([]string, 0)
	jwe := Jwe{}
	//1 生成头部
	header := NewHeader(ALG_RSA1_5, ENC_A128CBC_HS256)
	jsonHeader, err := utils.JsonEncode(header)
	args = append(args, jsonHeader)
	fmt.Println(jsonHeader, err)

	//2 加密密钥
	rsa := rsa.EncryptionMethodRSA{}
	publickey, _ := rsa.GetPublicKey()
	//	generateKey(1024)
	key, RasKey := jwe.GetEncryptedKey(header, 16, publickey)
	args = append(args, string(RasKey))
	fmt.Println("key:", string(key))
	fmt.Println("RsaKey:", string(RasKey))

	//3 4
	plant := "gaoqiankun"
	cipher, Iv := jwe.GetCipherText(header, []byte(plant), key)
	args = append(args, string(Iv))
	args = append(args, string(cipher))

	fmt.Println("cipher  Iv:", string(cipher), string(Iv))
	//5 得到数字证书
	Atag := jwe.GetAuthenticationTag(header, []string{string(RasKey), string(Iv), string(cipher)} /*[]byte{}*/, key)
	fmt.Println("Atag:", string(Atag))
	args = append(args, string(Atag))
	//6 得到jwe
	jw := jwe.GetJWE(args)
	fmt.Println("jwe:", jw)
}
