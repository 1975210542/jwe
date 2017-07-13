package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

type EncryptionMethodAES struct {
	Name string
}

var (
	EncryptionMethodAES256 *EncryptionMethodAES
)

func init() {
	// RS256
	EncryptionMethodAES256 = &EncryptionMethodAES{"A128CBC-HS256"}
	RegisterSigningMethod(EncryptionMethodAES256.GetName(), func() EncryptionMethod {
		return EncryptionMethodAES256
	})

}

func (e *EncryptionMethodAES) GetName() string {
	return e.Name
}

func (e *EncryptionMethodAES) Encrypt(plantText []byte, key interface{}) ([]byte, error) {
	fmt.Println("AES 加密！！！")
	return aesEncrypt(plantText, key)
}

func (e *EncryptionMethodAES) Decrypt(cipherText []byte, key interface{}) ([]byte, error) {
	fmt.Println("AES 解密！！！")
	return aesDecrypt(cipherText, key)
}

func (e *EncryptionMethodAES) GetKey(size int) []byte {
	return getAesKey(size)
}

/*****************************************/
func aesEncrypt(plantText []byte, key interface{}) ([]byte, error) {
	fmt.Println("AES 第二次调用")
	keyBytes := key.([]byte)
	block, err := aes.NewCipher(keyBytes) //选择加密算法
	if err != nil {
		return nil, err
	}
	plantText = pKCS7Padding(plantText, block.BlockSize())

	blockModel := cipher.NewCBCEncrypter(block, keyBytes)

	ciphertext := make([]byte, len(plantText))

	blockModel.CryptBlocks(ciphertext, plantText)
	return ciphertext, nil
}

func pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func getAesKey(size int) []byte {
	key := GenerateRandString(size)
	return key
}

//解密
func aesDecrypt(ciphertext []byte, key interface{}) ([]byte, error) {
	//	keyBytes := []byte(key)
	keyBytes := key.([]byte)
	fmt.Println("keyBytes:", keyBytes)
	block, err := aes.NewCipher(keyBytes) //选择加密算法
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, keyBytes)
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = pKCS7UnPadding(plantText, block.BlockSize())
	return plantText, nil
}

func pKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}
