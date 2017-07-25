package main

import (
	_ "jwe/aes"
	_ "jwe/hma"
	"jwe/methodit"
	_ "jwe/rsa"
	"jwe/utils"
	"log"
	"strings"
)

type Jwe struct{}

func (j *Jwe) GetEncryptedKey(header Header, size int, key interface{}) (aesKey, RsaKey []byte) {
	return getEncryptedKey(header, size, key)
}

func (j *Jwe) GetCipherText(header Header, plant []byte, key interface{}) (ciphertext, IV []byte) {
	return getCipherText(header, plant, key)
}

func (j *Jwe) GetAuthenticationTag(header Header, args []string, key interface{}) (Atag []byte) {
	return getAuthenticationTag(header, args, key)
}

func (j *Jwe) GetJWE(args []string) string {
	return getJWE(args)
}

func getEncryptedKey(header Header, size int, key interface{}) (aeskey, RsaKey []byte) {

	key = key.([]byte)
	alg := header.Alg
	method := methodit.GetSigningMethod(alg)
	aesKey := utils.GenerateRandString(size)

	RsaKey, err := method.Encrypt(aesKey, key)
	if err != nil {
		log.Println("加密出错了！！！", err)
	}
	return aesKey, RsaKey

}

func getCipherText(header Header, plant []byte, key interface{}) (ciphertext, IV []byte) {

	Key := key.([]byte)
	IV = []byte(Key)
	alg := strings.Split(header.Enc, "-")
	log.Print("alg:", alg, alg[0])
	method := methodit.GetSigningMethod(alg[0])
	ciphertext, err := method.Encrypt(plant, Key)
	if err != nil {
		log.Println("加密出错了", err)
	}
	return
}

func getAuthenticationTag(header Header, args []string, key interface{}) (Atag []byte) {

	alg := strings.Split(header.Enc, "-")
	tag := strings.Join(args, ".")
	log.Println("Alg:", alg, alg[1])
	method := methodit.GetSigningMethod(alg[1])
	Atag, err := method.Encrypt([]byte(tag), key)
	if err != nil {
		log.Println("加密出错了", err)
	}
	return
}
func getJWE(args []string) string {
	//	step 6. 拼接以及序列号数据，得到JWE Object 把以上5个步骤的数据进行Base64UrlEncode，然后按照顺序拼接，用"."分割，得到最后的数据。
	var arg []string
	for _, iterm := range args {
		arg = append(arg, utils.Base64Encode(iterm))
		log.Println("iterm:", iterm)
	}
	return strings.Join(arg, ".")

}

//func JweDecryp(jwe string) ([]byte, error) {
//	parts := strings.Split(jwe, ".")
//	if len(parts) != 5 {
//		return nil, errors.New("Wrong number of parts")
//	}

//	//decode jwe Header
//	var header Header
//	jso, err := Base64Decode(parts[0])
//	if err != nil {
//		log.Println("Header Decode faile!!!")
//	}

//	err = JsonDecode(string(jso), &header)
//	if err != nil {
//		log.Println("Json To Header faile!!")
//	}
//	log.Println(header.Alg, "   ", header.Enc)

//	//decode jwe key
//	rsa := EncryptionMethodRSA{}
//	RasKey, err := Base64Decode(parts[1])
//	key, err := rsa.Decrypt([]byte(RasKey))
//	if err != nil {
//		log.Println("Decode key faile")
//	}

//	log.Println("key:", key)
//	return nil, err
//}
