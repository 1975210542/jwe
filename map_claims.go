package main

import (
	//	"errors"
	"fmt"
	"time"
)

type MapClaims map[string]interface{}

func (mc MapClaims) Verif() {

	now := time.Now().Unix()
	if mc.VerifyExp(now, false) == false {

		exp := mc["exp"].(int64)
		subTime := time.Unix(now, 0).Sub(time.Unix(exp, 0))
		fmt.Println("The Token has expired by ", subTime)
	}

	if mc.VerifyIat(now, false) == false {
		fmt.Println("Use Token Before IssueAt")
	}

	if mc.VerifyNbf(now, false) == false {
		fmt.Println("The token is not yet valid")
	}

}

//Verify Exp
func (mc MapClaims) VerifyExp(now int64, rea bool) bool {
	exp := mc["exp"].(int64)
	return verifyExp(exp, now, rea)
}

//Verify Iat
func (mc MapClaims) VerifyIat(now int64, rea bool) bool {
	iat := mc["iat"].(int64)
	return verifyIat(iat, now, rea)
}

//Verify nbf
func (mc MapClaims) VerifyNbf(now int64, rea bool) bool {
	nbf := mc["nbf"].(int64)
	return verifyNbf(nbf, now, rea)
}

// Verify Aud
func (mc MapClaims) VerifyAud(cmp string, rea bool) bool {
	aud := mc["aud"].(string)
	return verifyAud(aud, cmp, rea)
}

//Verify Iss
func (mc MapClaims) VerifyIss(cmp string, rea bool) bool {
	iss := mc["iss"].(string)
	return verifyIss(iss, cmp, rea)
}
