package main

import (
	"crypto/subtle"
	//	"errors"
	"fmt"
	"time"
)

type Claims interface {
	Verif()
}

//iss：token签发者
//exp：token过期时间戳
//sub：token面向的用户/token的主题
//aud：token接收方
//iat：签发时间
//nbf:“Not before”，JWT不能接受处理的时间
//jti: JWT ID claim，为JWT提供唯一的标识符
type StandarClaims struct {
	Id        string `json:"jti,omitempty"`
	Audience  string `json:"aud,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	Subject   string `json:"sub,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	IssueAt   int64  `json:"iat,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
}

//验证token  的过期时间，签发时间和jwt不能就收处理的时间
func (sc *StandarClaims) Verif() {

	now := time.Now().Unix()
	if sc.VerifyExp(now, false) == false {
		subTime := time.Unix(now, 0).Sub(time.Unix(sc.ExpiresAt, 0))
		fmt.Println("The Token has expired by ", subTime)
	}

	if sc.VerifyIat(now, false) == false {
		fmt.Println("Use Token Before IssueAt")
	}

	if sc.VerifyNbf(now, false) == false {
		fmt.Println("The token is not yet valid")
	}

}

//Verify Exp
func (sc *StandarClaims) VerifyExp(now int64, rea bool) bool {
	return verifyExp(sc.ExpiresAt, now, rea)
}

//Verify Iat
func (sc *StandarClaims) VerifyIat(now int64, rea bool) bool {
	return verifyIat(sc.IssueAt, now, rea)
}

//Verify nbf
func (sc *StandarClaims) VerifyNbf(now int64, rea bool) bool {
	return verifyNbf(sc.NotBefore, now, rea)
}

// Verify Aud
func (sc *StandarClaims) VerifyAud(cmp string, rea bool) bool {
	return verifyAud(sc.Audience, cmp, rea)
}

//Verify Iss
func (sc *StandarClaims) VerifyIss(cmp string, rea bool) bool {
	return verifyIss(sc.Issuer, cmp, rea)
}

func verifyIss(iss, cmp string, rea bool) bool {
	if iss == "" {
		return !rea
	}
	if subtle.ConstantTimeCompare([]byte(iss), []byte(cmp)) == 0 {
		return false
	}
	return true
}

func verifyAud(aud, cmp string, rea bool) bool {
	if aud == "" {
		return !rea
	}

	if subtle.ConstantTimeCompare([]byte(aud), []byte(cmp)) == 0 {
		return false
	}

	return true
}
func verifyIat(iat, now int64, rea bool) bool {
	if iat == 0 {
		return !rea
	}
	return now >= iat
}

func verifyNbf(nbf, now int64, rea bool) bool {
	if nbf == 0 {
		return !rea
	}
	return nbf >= now
}

func verifyExp(exp, now int64, rea bool) bool {
	if exp == 0 {
		return !rea
	}
	return now <= exp
}
