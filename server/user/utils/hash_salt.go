package utils

import (
	"crypto/sha256"
	"fmt"
)

// HashSalt 生产环境需要替换以下字符串
var HashSalt = "123456adbcdef"

// GetHash Create the salted hash
func GetHash(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	hash.Write([]byte(HashSalt))
	hashBytes := hash.Sum(nil)
	return fmt.Sprintf("%x", hashBytes)
}
