package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_GetRsaKey(t *testing.T) {
	rsa := EncryptionMethodRSA{}
	if _, err := rsa.GetPrivateKey(); err != nil {
		t.Error("Did not get the privateKey !!!")
	} else {
		t.Log("Get the privateKey  Successfully!")
	}

	if _, err := rsa.GetPublicKey(); err != nil {
		t.Error("Did not get the publicKey ！！！")
	} else {
		t.Log("Get the publicKey Successfully!")
	}
}

func Test_EncrptyAndDecrpty(t *testing.T) {
	rsa := EncryptionMethodRSA{}
	plant := "China!"
	var publicKey, privateKey, cipher, pl []byte
	var err error

	if publicKey, err = rsa.GetPublicKey(); err != nil {
		t.Error("Did not get the publicKey !!")
	}

	if cipher, err = rsa.Encrypt([]byte(plant), publicKey); err != nil {
		t.Error("Encrption Failed!!")
	} else {
		t.Log("Encrption Successfully!")
	}
	fmt.Println("cipher:", cipher)
	if privateKey, err = rsa.GetPrivateKey(); err != nil {
		t.Error("Did not get the privatekey!!")
	}

	if pl, err = rsa.Decrypt(cipher, privateKey); err != nil || strings.Compare(string(pl), plant) != 0 {
		t.Error("Decrpty Failed!")
	} else {
		t.Log("Decrption Successfully!")
	}
	fmt.Println("original:", string(pl))
}
