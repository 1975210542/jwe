package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_AesEncrpty(t *testing.T) {
	aes := EncryptionMethodAES{}
	key := GenerateRandString(16)
	fmt.Println("woyaodekey:", key)
	plant := []byte("China!")
	var cipher, original []byte
	var err error

	if cipher, err = aes.Encrypt(plant, key); err != nil {
		t.Error("AES Encrpty Failed!!!")
	} else {
		t.Log("AES Encrpty Successfully!!")
	}
	fmt.Println("AEScipher:", cipher)

	if original, err = aes.Decrypt(cipher, key); strings.Compare(string(original), string(plant)) != 0 {
		t.Error("AES Decrpty Failed!!!")
	} else {
		fmt.Println("AES解密：cipher:", cipher)
		t.Log("AES Decrpty Successfuly!!!")
	}
}
