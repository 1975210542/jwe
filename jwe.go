package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"strings"
)

type Jwe struct{}

func (j *Jwe) GetEncryptedKey(header Header, size int) (key []byte, RsaKey string) {
	return getEncryptedKey(header, size)
}

func (j *Jwe) GetCipherText(header Header, key []byte, plant string) (ciphertext, IV []byte) {
	return getCipherText(header, key, plant)
}
func (j *Jwe) GetAuthenticationTag(header Header, args []string) (Atag string) {
	return getAuthenticationTag(header, args)
}

func (j *Jwe) GetJWE(args []string) string {
	return getJWE(args)
}

func getEncryptedKey(header Header, size int) (key []byte, RsaKey string) {
	//	step 2. 生成密钥并且加密密钥，得到Encrypted Key//随机生成一组AES的Key，长度为128，然后用RSA -pkcs1-v1_5进行加密

	switch header.Alg {
	case ALG_RSA1_5:
		aes := EncryptionMethodAES{}
		key = aes.GetKey(size)
		RsaKey = rsaencrypt(string(key))
		log.Println("Base64RsaKey:", RsaKey)
	default:
		log.Println("Unsupported ALG keytype")
	}
	return

}

func getCipherText(header Header, key []byte, plant string) (ciphertext, IV []byte) {
	//	step 3. 生成向量数据Iv，得到Initialization Vector IV和密钥一样
	switch header.Enc {
	case ENC_A128CBC_HS256:
		IV = []byte(key)
		log.Println("IV:", IV)

		//step 4. 加密原始报文，得到Cipher text  利用step2的密钥 和 step3的向量数据，用AES(CBC)128加密
		ciphertext = encryptRawMessage(plant, key)
		log.Println("密文:", ciphertext)
	default:
		log.Println("Unsupported enc type")
	}
	return
}

func getAuthenticationTag(header Header, args []string) (Atag string) {

	//step 5. 生成认证码，得到Authentication Tag
	//把step2的加密密钥、step3的向量、step4的密文 进行拼接，然后用HMAC-SHA256 算法进行签名
	switch header.Enc {
	case ENC_A128CBC_HS256:
		tag := strings.Join(args, ".")
		log.Println("tag:", tag)
		Atag = generateAuthenticationTag(tag)
		log.Println("Atag:", Atag)
	default:
		log.Println("Unsupported enc type")
	}
	return
}
func getJWE(args []string) string {
	//	step 6. 拼接以及序列号数据，得到JWE Object 把以上5个步骤的数据进行Base64UrlEncode，然后按照顺序拼接，用"."分割，得到最后的数据。
	var arg []string
	for _, iterm := range args {
		arg = append(arg, Base64Encode(iterm))
		log.Println("iterm:", iterm)
	}
	return strings.Join(arg, ".")

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
	rsa := EncryptionMethodRSA{}
	RasKey, err := Base64Decode(parts[1])
	key, err := rsa.Decrypt([]byte(RasKey))
	if err != nil {
		log.Println("Decode key faile")
	}

	log.Println("key:", key)
	return nil, err
}

//生成头部
func GenerateJweHeader(header Header) string {

	var jh string
	if json, err := JsonEncode(header); err == nil {
		jh = Base64Encode(json)
	}
	return jh
}

//对生成的密钥进行RSA加密
func rsaencrypt(key string) string {
	rsa := EncryptionMethodRSA{}
	rsa.GenerateKey(1024)
	data, err := rsa.Encrypt([]byte(key)) //RSA加密
	if err != nil {
		log.Println(err)
	}
	return string(data)
}

//原始加密原始报文，得到Cipher text
func encryptRawMessage(plant string, key []byte) []byte {
	//	key := getKey()
	aes := EncryptionMethodAES{}
	ciphertext, err := aes.Encrypt([]byte(plant), key)
	if err != nil {
		log.Println("err:", err)
	}
	return ciphertext
}

//生成认证码，得到Authentication Tag
func generateAuthenticationTag(s string) string {

	mac := hmac.New(sha256.New, []byte(""))
	mac.Write([]byte(s))
	return hex.EncodeToString(mac.Sum(nil))
}
