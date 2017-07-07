package main

import (
	//	"crypto/rand"
	//	"crypto/rsa"
	//	"crypto/x509"
	//	"encoding/pem"
	"fmt"
	//	"log"
)

func main() {
	//	GenerateKey(1024)
	//	data, err := RsaEncrypt([]byte("qiankun")) //RSA加密
	//	if err != nil {
	//		panic(err)
	//	}
	key := GenerateEncryptedKey(16)
	fmt.Println("key:", key)
	RsaKey := RSAEncrypt(key)
	fmt.Println("RSA加密", RsaKey)

	origData, err := RsaDecrypt([]byte(RsaKey)) //RSA解密
	if err != nil {
		panic(err)
	}
	fmt.Println("RSA解密", string(origData))

}
