package main

import (
	//	"crypto/rand"
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
	plant := "gaoqiankun"
	jwe1 := GenerateJWE(plant, ALG_RSA1_5, ENC_A128CBC_HS256, 16)
	fmt.Println("jwe1:", jwe1)

	//	jwe := "eyJhbGciOiJSU0ExXzUiLCJlbmMiOiJBMTI4Q0JDLUhTMjU2In0.ep97JdQldzCVdGxSB_r4GynGrH2t8Au_T0LhpHQlZIn_SXVL6MyrJdKsTnFwzS-TZlrne94ZwlQeiwWrdJfu5QicRavQ-pgGDBe5vadfaLgYuAHqI70hY8AdY5n2TVOU7Q-ZLqO41xOv40JV0h-pZxe0fysWcNCD4D9SMXgZx3U.enZrNnRpemgyaW9rYTFxbg.2ypqAprYGAaSf7T-pHRfUg.ODE2OTVmNjZhZmU1OTA4NWI2MDc5MDgzNTkxNTZiNmE1YmFhNTZmYjljNmUyYmUwNmE4ZTJiNTFlZWU3ZGIzZQ"
	//	str, err := JweDecryp(jwe)
	//	fmt.Println(str, err)
}
