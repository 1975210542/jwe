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
	var args []string
	jwe := Jwe{}
	//1
	header := NewHeader(ALG_RSA1_5, ENC_A128CBC_HS256)

	jHeader, _ := JsonEncode(header)
	args = append(args, jHeader)
	fmt.Println("jHeader:", jHeader)
	//2
	key, RsaKey := jwe.GetEncryptedKey(header, 16)
	args = append(args, RsaKey)
	fmt.Println(RsaKey)
	//3 4
	plant := "gaoqiankun"
	cipher, Iv := jwe.GetCipherText(header, key, plant)
	fmt.Println("Cipher!!!!:", cipher)
	args = append(args, string(cipher))
	args = append(args, string(Iv))
	fmt.Println("len:", len(args))
	//5
	Atag := jwe.GetAuthenticationTag(header, []string{args[1], args[2], args[3]})
	fmt.Println("Atag:", Atag)
	args = append(args, Atag)
	//6
	jw := jwe.GetJWE(args)
	fmt.Println("jwe:", jw)
}
