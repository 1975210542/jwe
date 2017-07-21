package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"jwe/methodit"
	"log"
	"os"
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
	methodit.RegisterSigningMethod(EncryptionMethodRSA256.GetName(), func() methodit.EncryptionMethod {
		return EncryptionMethodRSA256
	})

}

func (e *EncryptionMethodRSA) GetName() string {
	return e.Name
}

func (e *EncryptionMethodRSA) Encrypt(plantText []byte, key interface{}) ([]byte, error) {
	//	fmt.Println("RSA jiami")
	return rsaEncrypt(plantText, key)
}

func (e *EncryptionMethodRSA) Decrypt(cipherText []byte, key interface{}) ([]byte, error) {
	//	fmt.Println("RSA jiemi")
	return rsaDecrypt(cipherText, key)
}

//生成密钥对
func (e *EncryptionMethodRSA) GenerateKey(bits int) {
	generateKey(bits)
}

//private publicKey
func (e *EncryptionMethodRSA) GetPublicKey() (publicKey []byte, err error) {

	if publicKey, err = ioutil.ReadFile("test/publicKey.pem"); err != nil {
		return []byte{}, err
	}
	return publicKey, nil
}

//get privateKey
func (e *EncryptionMethodRSA) GetPrivateKey() (privateKey []byte, err error) {

	if privateKey, err = ioutil.ReadFile("test/privateKey.pem"); err != nil {
		return []byte{}, err
	}
	return privateKey, nil
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

	file, err := os.Create("test/privateKey.pem")
	if err != nil {
		fmt.Println("Failed to Create privateKey's file!!! ", err)
	}
	err = pem.Encode(file, block)
	if err != nil {
		fmt.Println("Failed to Encode private key!!!", err)
	}

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

	file, err = os.Create("test/publicKey.pem")
	if err != nil {
		fmt.Println("Failed to create publicKey's file")
	}

	err = pem.Encode(file, block)
	if err != nil {
		fmt.Println("Failed to Encode public key!!!")
	}

}

// 加密
func rsaEncrypt(origData []byte, key interface{}) ([]byte, error) {
	//	fmt.Println("RSA 第二次调用")
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
