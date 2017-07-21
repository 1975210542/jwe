package methodit

import (
	"fmt"
)

var encryptionMethod = map[string]func() EncryptionMethod{}

type EncryptionMethod interface {
	Encrypt(plantText []byte, key interface{}) ([]byte, error)
	Decrypt(cipherText []byte, key interface{}) ([]byte, error)
	GetName() string
}

func RegisterSigningMethod(alg string, f func() EncryptionMethod) {
	fmt.Println("注册方法！！！")
	encryptionMethod[alg] = f
}

func GetSigningMethod(alg string) (method EncryptionMethod) {
	fmt.Println("得到方法！！！")
	if methodF, ok := encryptionMethod[alg]; ok {
		method = methodF()
	}
	return
}
