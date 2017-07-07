package main

import (
	"log"
	//	"strings"
)

//生成头部
func GenerateJweHeader(header Header) string {

	var jh string
	if json, err := JsonEncode(header); err == nil {
		jh = Base64Encode(json)
	}
	return jh
}

//生成密钥
func GenerateEncryptedKey(size int) string {
	return GenerateRandString(size)
}

//对生成的密钥进行RSA加密
func RSAEncrypt(key string) string {
	GenerateKey(1024)
	data, err := RsaEncrypt([]byte(key)) //RSA加密
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

func GenerateJWE() {
	//选择算法，生成JWEHeader
	//密钥加密算法为RSA,padding方式为：pkcs1-v1_5
	//原始报文的加密算法为AES128，提供商为CBC。HS256：签名算法为HMAC With SHA256
	//JWEHeader为：{"alg":"RSA1_5","enc":"A128CBC-HS256"}
	header := NewHeader()
	header.SetHeader(ALG_RSA1_5, ENC_A128CBC_HS256)
	jweheader := GenerateJweHeader(header)
	log.Println("jweHeader:", jweheader)
	//step 2. 生成密钥并且加密密钥，得到Encrypted Key//随机生成一组AES的Key，长度为128，然后用RSA -pkcs1-v1_5进行加密。
	//注：此处用的RSA的公钥
	key := GenerateEncryptedKey(16)
	RsaKey := RSAEncrypt(key)
	Base64RsaKey := Base64Encode(RsaKey)
	//	step 3. 生成向量数据，得到Initialization Vector

}
