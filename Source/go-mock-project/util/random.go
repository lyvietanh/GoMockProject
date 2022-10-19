package util

import (
	"math/rand"
	"strings"
	"time"
)

const ALPHA_NUMERIC = "abcdefghijklmnopqrstuvwyzABCDEFGHIJKLMNOPQRSTUVWYZ0123456789"

var countryCodes = []string{"VN", "US", "JP", "KR", "EU", "GB"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateRandomString(length int) string {
	var sb strings.Builder
	n := len(ALPHA_NUMERIC)
	for i := 0; i < length; i++ {
		c := ALPHA_NUMERIC[rand.Intn(n)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func GenerateRandomCountryCode() string {
	n := len(countryCodes)
	result := countryCodes[rand.Intn(n)]
	return result
}
