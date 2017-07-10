package main

import (
	//	"crypto/rand"
	//	"crypto/rsa"
	//	"crypto/x509"
	//	"encoding/pem"
	"fmt"
	//	"strings"
	//	"log"
)

func main() {
	plant := "gaoqiankun"

	jwe := GenerateJWE(plant)
	//	strings.Join()
	fmt.Println("jwe:", jwe)
	//	"eyJhbGciOiJSU0ExXzUiLCJlbmMiOiJBMTI4Q0JDLUhTMjU2In0.
	//	K2wrP1uv0slZO0xP6eeqafuiZA1gxRaB0iJnJFT3oLqGcrkMwZkq7pKFxz8r6Kh5jeXe9twbG3qjGIO-x-T0dISLQ1h4_5NN-UgFB99zXhDbvgiJR54G3SK5ZYMQhofbz9TWnUeksRRQzTYWJWyp41eeyWtxf_kjmG8UKBQFeXI.
	//	OTM3Z2FmdnN2djg1NWljdw.
	//	Mrn3Zc9G8Naf9vYG7KQDtQ.
	//	ODA3MTYzZDQxNzRmYzg4NzM2Yzk2OThiMTRhNTYxNjlhYTE4ZTk5Yzk1Y2UxMDkyZTQxM2M1OWEyYmM0NGY1OQ"
}
