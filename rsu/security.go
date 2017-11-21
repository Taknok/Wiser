package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func ComputeHmac256(message string, secret string) []byte {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return (h.Sum(nil))
}

// CheckMAC reports whether messageMAC is a valid HMAC tag for message.
func CheckMAC(message, messageMAC  []byte, secret string) bool {
	key := []byte(secret)
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

func main() {
	secret :="jesuislaclesecrete"
	idVehicule := "123"
	idRsu := "456"
	date := "789"

	message := idVehicule+idRsu+date
	fmt.Println(message)

	messageMAC := ComputeHmac256(message, secret)
	fmt.Println("message : ",message)
	fmt.Println("messageMAC : ", messageMAC)

	if CheckMAC([]byte(message), messageMAC, secret)== true {
		fmt.Println("MAC --> OK")
	} else {
		fmt.Println("MAC --> ERROR")
	}
}