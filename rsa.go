package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
)

var (
	PRIVATEKEY []byte
	PUBLICKEY  []byte
)

type EncryptionMethodRSA struct {
	Name string
}

var (
	EncryptionMethodRSA256 *EncryptionMethodRSA
)

func init() {
	// RS256
	EncryptionMethodRSA256 = &EncryptionMethodRSA{"RSA1_5"}
	RegisterSigningMethod(EncryptionMethodRSA256.GetName(), func() EncryptionMethod {
		return EncryptionMethodRSA256
	})

}

func (e *EncryptionMethodRSA) GetName() string {
	return e.Name
}
func (e *EncryptionMethodRSA) GenerateKey(bits int) {
	generateKey(bits)
}

func (e *EncryptionMethodRSA) Encrypt(plantText []byte, key interface{}) ([]byte, error) {
	fmt.Println("RSA jiami")
	key = PUBLICKEY
	return rsaEncrypt(plantText, key)
}

func (e *EncryptionMethodRSA) Decrypt(cipherText []byte, key interface{}) ([]byte, error) {
	fmt.Println("RSA jiemi")
	key = PRIVATEKEY
	return rsaDecrypt(cipherText, key)
}

func generateKey(size int) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
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

}

// 加密
func rsaEncrypt(origData []byte, key interface{}) ([]byte, error) {
	fmt.Println("RSA 第二次调用")
	publicKey := key.([]byte)
	block, _ := pem.Decode(publicKey) //将密钥解析成公钥实例
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(origData)) //RSA算法加密
}

// 解密
func rsaDecrypt(ciphertext []byte, key interface{}) ([]byte, error) {
	privateKey := key.([]byte)
	block, _ := pem.Decode(privateKey) //将密钥解析成私钥实例
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, []byte(ciphertext)) //RSA算法解密
}
