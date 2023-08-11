package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Convert string to Md5
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
