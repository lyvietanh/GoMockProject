package util

import (
	"crypto/sha256"
	"fmt"
	"log"
)

func GenerateSha256(text string) string {
	log.Println(text)
	hash := sha256.New()
	hash.Write([]byte(text))
	result := fmt.Sprintf("%x", hash.Sum(nil))
	log.Println(result)
	return result
}

func GenerateSha256SecurePassword(email string, clearPassword string, passwordSalt string) string {
	return GenerateSha256(fmt.Sprintf("%v|%v|%v", email, clearPassword, passwordSalt))
}
