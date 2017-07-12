package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
)

var (
	PRIVATEKEY []byte
	PUBLICKEY  []byte
)

type EncryptionMethodRSA struct{}

func (e *EncryptionMethodRSA) GenerateKey(bits int) {
	generateKey(bits)
}

func (e *EncryptionMethodRSA) Encrypt(origData []byte) ([]byte, error) {
	return rsaEncrypt(origData)
}

func (e *EncryptionMethodRSA) Decrypt(ciphertext []byte) ([]byte, error) {
	return rsaDecrypt(ciphertext)
}

func generateKey(bits int) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Println(err)
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	PRIVATEKEY = pem.EncodeToMemory(block)

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Println(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	PUBLICKEY = pem.EncodeToMemory(block)
	//	log.Println(string(PRIVATEKEY))
	//	log.Println(string(PUBLICKEY))
}

// 加密
func rsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(PUBLICKEY) //将密钥解析成公钥实例
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData) //RSA算法加密
}

// 解密
func rsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(PRIVATEKEY) //将密钥解析成私钥实例
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext) //RSA算法解密
}
