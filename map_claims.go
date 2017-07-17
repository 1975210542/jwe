package main

import (
	"errors"
	"fmt"
)

type MapClaims map[string]interface{}

func (mc MapClaims) Verif() error {
	fmt.Println("MapClaims!!!")
	return errors.New("MapClaims errors!!!")
}
