package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"strings"
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
func GenerateEncryptedKey(size int) []byte {
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

//原始加密原始报文，得到Cipher text
func EncryptRawMessage(plant string, key []byte) []byte {
	//	key := getKey()
	ciphertext, err := Encrypt([]byte(plant), key)
	if err != nil {
		log.Println("err:", err)
	}
	return ciphertext
}

//生成认证码，得到Authentication Tag
func GetAuthenticationTag(s string) string {

	mac := hmac.New(sha256.New, []byte(""))
	mac.Write([]byte(s))
	return hex.EncodeToString(mac.Sum(nil))
}
func GenerateJWE(plant string) string {
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
	RsaKey := RSAEncrypt(string(key))
	Base64RsaKey := Base64Encode(RsaKey)
	log.Println("Base64RsaKey:", Base64RsaKey)

	//	step 3. 生成向量数据Iv，得到Initialization Vector IV和密钥一样
	IV := []byte(key)
	baseIv := Base64Encode(string(IV))
	log.Println(baseIv)

	//step 4. 加密原始报文，得到Cipher text  利用step2的密钥 和 step3的向量数据，用AES(CBC)128加密
	ciphertext := EncryptRawMessage(plant, key)
	baseciphertext := Base64Encode(string(ciphertext))
	log.Println("密文:", ciphertext)

	//step 5. 生成认证码，得到Authentication Tag
	//把step2的加密密钥、step3的向量、step4的密文 进行拼接，然后用HMAC-SHA256 算法进行签名
	tag := strings.Join([]string{RsaKey, string(IV), string(ciphertext)}, ".")
	log.Println("tag:", tag)
	Atag := GetAuthenticationTag(tag)
	baseAtag := Base64Encode(Atag)
	log.Println("Atag:", Atag)
	//	step 6. 拼接以及序列号数据，得到JWE Object 把以上5个步骤的数据进行Base64UrlEncode，然后按照顺序拼接，用"."分割，得到最后的数据。
	return strings.Join([]string{jweheader, Base64RsaKey, baseIv, baseciphertext, baseAtag}, ".")

}
