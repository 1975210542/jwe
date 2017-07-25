package hma

import (
	"testing"
)

func Test_Hmac(t *testing.T) {
	hmac := EncryptionMethodHMAC{}
	plant := []byte("China!")
	key := GenerateRandString(16)
	var cipher []byte
	var err error

	if cipher, err = hmac.Encrypt(plant, key); err != nil {
		t.Error("Hmac Encrpty Failed!!")
	} else {
		t.Log("Hamc Encrpty Successfuly!!")
	}

}
