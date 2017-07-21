package rsa

//import (
//	"testing"
//)

//func Benchmark_Encrpty(b *testing.B) {
//	rsa := EncryptionMethodRSA{}
//	plant := "China!"
//	//	var publicKey, privateKey, cipher, pl []byte
//	//	var err error

//	for i := 0; i < b.N; i++ {
//		publicKey, _ := rsa.GetPublicKey()
//		rsa.Encrypt([]byte(plant), publicKey)

//	}

//}

//func Benchmark_Decrpty(b *testing.B) {

//	b.StopTimer()

//	b.StartTimer()
//	rsa := EncryptionMethodRSA{}
//	cipher := []byte{12, 56, 78, 65, 1, 2, 3}

//	for i := 0; i < b.N; i++ {
//		privateKey, _ := rsa.GetPrivateKey()
//		rsa.Decrypt(cipher, privateKey)
//	}
//}
