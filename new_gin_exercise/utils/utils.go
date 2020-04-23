package utils

import (
	"math/rand"
	"time"
)

func RandString(n int) string {
	var letters = []byte("sdfsdfnidngoSDNDSOIFNIOnsdnfnsdflDSNGOIDfdgdsfgfdgSNIOFGNO")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
