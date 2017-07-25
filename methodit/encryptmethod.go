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
	fmt.Println("注册方法！！！", alg)
	encryptionMethod[alg] = f
}

func GetSigningMethod(alg string) (method EncryptionMethod) {

	if methodF, ok := encryptionMethod[alg]; ok {
		method = methodF()
		fmt.Println("得到方法！！！")
	}
	return
}
