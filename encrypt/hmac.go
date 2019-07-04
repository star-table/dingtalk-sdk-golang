package encrypt

import (
	"crypto/hmac"
	"crypto/sha256"
)

func SHA256(source string, secretKey string) []byte {
	key := []byte(secretKey)
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(source))
	return mac.Sum(nil)
}
