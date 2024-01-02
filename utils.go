package main

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strings"
)

func generateURLID(longURL string) string {
	uniqueID := generateRandomString(6)
	hashedURL := hashSHA256(longURL)[:6]
	return hashedURL + uniqueID
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result strings.Builder
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		result.WriteByte(charset[randomIndex])
	}
	return result.String()
}

func hashSHA256(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	hashInBytes := hash.Sum(nil)
	return hex.EncodeToString(hashInBytes)
}
