package main

import ( //	"crypto/rand"
	//	"crypto/rsa"
	//	"crypto/x509"
	//	"encoding/pem"
	"fmt"

	//	"encoding/base64"
	//	"strings"
	//	"log"
)

/*eyJhbGciOiJSU0ExXzUiLCJlbmMiOiJBMTI4Q0JDLUhTMjU2In0
.ep97JdQldzCVdGxSB_r4GynGrH2t8Au_T0LhpHQlZIn_SXVL6MyrJdKsTnFwzS-TZlrne94ZwlQeiwWrdJfu5QicRavQ-pgGDBe5vadfaLgYuAHqI70hY8AdY5n2TVOU7Q-ZLqO41xOv40JV0h-pZxe0fysWcNCD4D9SMXgZx3U
.enZrNnRpemgyaW9rYTFxbg
.2ypqAprYGAaSf7T-pHRfUg
.ODE2OTVmNjZhZmU1OTA4NWI2MDc5MDgzNTkxNTZiNmE1YmFhNTZmYjljNmUyYmUwNmE4ZTJiNTFlZWU3ZGIzZQ*/
func main() {
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
