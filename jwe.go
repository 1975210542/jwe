package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
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
func GenerateJWE(plant string, alg Algorithm, enc EncryptionMethod, size int) string {
	args := make([]string, 0)
	step5 := make([]string, 0)

	//选择算法，生成JWEHeader
	//密钥加密算法为RSA,padding方式为：pkcs1-v1_5
	//原始报文的加密算法为AES128，提供商为CBC。HS256：签名算法为HMAC With SHA256
	//JWEHeader为：{"alg":"RSA1_5","enc":"A128CBC-HS256"}
	header := NewHeader()
	header.SetHeader(alg, enc)
	jweheader := GenerateJweHeader(header)
	args = append(args, jweheader)
	log.Println("jweHeader:", jweheader)

	//step 2. 生成密钥并且加密密钥，得到Encrypted Key//随机生成一组AES的Key，长度为128，然后用RSA -pkcs1-v1_5进行加密。
	//注：此处用的RSA的公钥
	var key []byte
	switch alg {
	case ALG_RSA1_5:
		key = GenerateEncryptedKey(size)
		RsaKey := RSAEncrypt(string(key))

		step5 = append(step5, RsaKey)
		Base64RsaKey := Base64Encode(RsaKey)
		args = append(args, Base64RsaKey)
		log.Println("Base64RsaKey:", Base64RsaKey)
	default:
		log.Println("Unsupported ALG keytype")
	}

	//	step 3. 生成向量数据Iv，得到Initialization Vector IV和密钥一样
	switch enc {
	case ENC_A128CBC_HS256:
		IV := []byte(key)
		step5 = append(step5, string(IV))

		baseIv := Base64Encode(string(IV))
		args = append(args, baseIv)
		log.Println(baseIv)
	default:
		log.Println("Unsupported enc type")
	}

	//step 4. 加密原始报文，得到Cipher text  利用step2的密钥 和 step3的向量数据，用AES(CBC)128加密
	switch enc {
	case ENC_A128CBC_HS256:
		ciphertext := EncryptRawMessage(plant, key)
		step5 = append(step5, string(ciphertext))

		baseciphertext := Base64Encode(string(ciphertext))
		args = append(args, baseciphertext)
		log.Println("密文:", ciphertext)
	default:
		log.Println("Unsupported enc type")
	}

	//step 5. 生成认证码，得到Authentication Tag
	//把step2的加密密钥、step3的向量、step4的密文 进行拼接，然后用HMAC-SHA256 算法进行签名
	switch enc {
	case ENC_A128CBC_HS256:
		tag := strings.Join(step5, ".")
		log.Println("tag:", tag)
		Atag := GetAuthenticationTag(tag)
		baseAtag := Base64Encode(Atag)
		args = append(args, baseAtag)
		log.Println("Atag:", Atag)
	default:
		log.Println("Unsupported enc type")
	}

	//	step 6. 拼接以及序列号数据，得到JWE Object 把以上5个步骤的数据进行Base64UrlEncode，然后按照顺序拼接，用"."分割，得到最后的数据。
	return strings.Join(args, ".")

}

func JweDecryp(jwe string) ([]byte, error) {
	parts := strings.Split(jwe, ".")
	if len(parts) != 5 {
		return nil, errors.New("Wrong number of parts")
	}

	//decode jwe Header
	var header Header
	jso, err := Base64Decode(parts[0])
	if err != nil {
		log.Println("Header Decode faile!!!")
	}

	err = JsonDecode(string(jso), &header)
	if err != nil {
		log.Println("Json To Header faile!!")
	}
	log.Println(header.Alg, "   ", header.Enc)

	//decode jwe key
	RasKey, err := Base64Decode(parts[1])
	key, err := RsaDecrypt([]byte(RasKey))
	if err != nil {
		log.Println("Decode key faile")
	}

	log.Println("key:", key)
	return nil, err
}
