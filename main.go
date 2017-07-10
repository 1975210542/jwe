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
}
