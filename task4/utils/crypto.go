package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashPassword 使用 SHA-256 加密密码
func HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}
