package utils

import (
	"math/rand"
	//	"fmt"
	"strings"
	"time"
)

func Merge(arg []string) string {
	return strings.Join(arg, ".")
}

func GenerateRandString(size int) []byte {

	bytes := "0123456789abcdefghijklmnopqrstuvwxyz"
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return result
}
