package password 

import (
	"encoding/base64"
	"math/rand"
	"crypto/sha256"
	"time"
)

const randomLength = 16

func GenerateSalt(length int) string {
	var salt []byte
	var asciiPad int64

	if length == 0 {
		lenght = randomLength
	}

	asciiPad = 32

	for i := 0; i < length; i ++ {
		salt = append(salt, byte(rand.Int63n(94) + asciiPad))
	} 

	return string(salt)
}

func GenerateHash(salt, password string) string {
	var hash string
	fullString := salt + password
	sha := sha256.New()
	sha.Write([]byte(fullString))
	hash := base64.URLEncoding.EncodeToString(sha.Sum(nil))

	return hash
}

func ReturePassword(password string) (string, string) {
	rand.Seed(time.Now().UTC().UnixNano())
	salt := GenerateSalt(0)
	hash := GenerateHash(salt, password)	
	return salt, hash
}
