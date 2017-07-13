package main

import ( //	"crypto/rand"
	//	"crypto/rsa"
	//	"crypto/x509"
	//	"encoding/pem"
	"fmt"
	"io/ioutil"
	//	"os"

	//	"encoding/base64"
	//	"strings"
	//	"log"
)

func main() {
	key := GenerateRandString(16)
	plant := "gaoqiankun"
	cipher, err := hmacEncrpt([]byte(plant), key)
	fmt.Println("cipher:", string(cipher), err)
}

func testrsa() {
	publicKey, _ := ioutil.ReadFile("test/publicKey.pem")
	fmt.Println("publicKey:", string(publicKey))
	plant := "gaoqinakun"
	cipher, err := rsaEncrypt([]byte(plant), string(publicKey))
	fmt.Println("cipher:", string(cipher), err)

	privateKey, _ := ioutil.ReadFile("test/privateKey.pem")
	clear, err := rsaDecrypt(cipher, string(privateKey))
	fmt.Println("clear:", string(clear), err)
}
func TestJwe() {
	args := make([]string, 0)
	jwe := Jwe{}
	//1 生成头部
	header := NewHeader(ALG_RSA1_5, ENC_A128CBC_HS256)
	jsonHeader, err := JsonEncode(header)
	args = append(args, jsonHeader)
	fmt.Println(jsonHeader, err)

	//2 加密密钥
	generateKey(1024)
	key, RasKey := jwe.GetEncryptedKey(header, 16, []byte{})
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
	Atag := jwe.GetAuthenticationTag(header, []string{string(RasKey), string(Iv), string(cipher)}, []byte{})
	fmt.Println("Atag:", string(Atag))
	args = append(args, string(Atag))
	//6 得到jwe
	jw := jwe.GetJWE(args)
	fmt.Println("jwe:", jw)
}
