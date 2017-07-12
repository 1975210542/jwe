package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

type EncryptionMethodAES struct{}

func (e *EncryptionMethodAES) Encrypt(plantText, key []byte) ([]byte, error) {
	return aesEncrypt(plantText, key)
}

func (e *EncryptionMethodAES) Decrypt(ciphertext, key []byte) ([]byte, error) {
	return aesDecrypt(ciphertext, key)
}

func (e *EncryptionMethodAES) GetKey(size int) []byte {
	return getAesKey(size)
}

func aesEncrypt(plantText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key) //选择加密算法
	if err != nil {
		return nil, err
	}
	plantText = pKCS7Padding(plantText, block.BlockSize())

	blockModel := cipher.NewCBCEncrypter(block, key)

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
func aesDecrypt(ciphertext, key []byte) ([]byte, error) {
	//	keyBytes := []byte(key)
	keyBytes := key
	fmt.Println("keyBytes:", keyBytes)
	block, err := aes.NewCipher(key) //选择加密算法
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, key)
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
