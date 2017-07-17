package main

import (
	"errors"
	"fmt"
)

type Claims interface {
	Verif() error
}

type StandarClaims struct {
	Audience string `json:"aud,omitempty"`
}

func (sc *StandarClaims) Verif() error {
	fmt.Println("StandarClaims!!")
	return errors.New("StanderClaims error!!!")
}
