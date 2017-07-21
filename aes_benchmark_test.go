package main

import (
	"testing"
)

func Benchmark_AesEncrpty(b *testing.B) {

	aes := EncryptionMethodAES{}
	plant := []byte("China!")
	key := []byte{104, 55, 114, 104, 104, 52, 108, 53, 54, 50, 110, 107, 115, 110, 111, 54}
	for i := 0; i < b.N; i++ {

		aes.Encrypt(plant, key)
	}
	cipher := []byte{65, 13, 37, 26, 168, 216, 248, 132, 47, 78, 8, 240, 247, 6, 130, 131}
	//	key := []byte{104, 55, 114, 104, 104, 52, 108, 53, 54, 50, 110, 107, 115, 110, 111, 54}
	for i := 0; i < b.N; i++ {
		aes.Decrypt(cipher, key)
	}
}

//func Benchmark_AesDencrpty(b *testing.B) {
//	aes := EncryptionMethodAES{}
//	cipher := []byte{65, 13, 37, 26, 168, 216, 248, 132, 47, 78, 8, 240, 247, 6, 130, 131}
//	key := []byte{104, 55, 114, 104, 104, 52, 108, 53, 54, 50, 110, 107, 115, 110, 111, 54}
//	for i := 0; i < b.N; i++ {

//		aes.Decrypt(cipher, key)
//	}
//}
