package common

import (
	"math/rand"
	"time"
)

const (
	VERIFICATIONCODELENGTH = 6

	VERIFICATIONCODEEXPIRATION = 60 * 5

	TYPEREGISTER = "register"
	TYPEPWD      = "pwd"
	TYPECRITICAL = "critical"
)

func GenerateVerificationCode() string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < VERIFICATIONCODELENGTH; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
