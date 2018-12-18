package tripismapiutilities

import (
	"crypto/md5"
	"encoding/hex"
)

// CalculateMD5 generates MD5 hash for a string
func CalculateMD5(input string) string {
	md5hash := md5.Sum([]byte(input))
	return hex.EncodeToString(md5hash[:])
}

// FileMD5 generates MD5 hash for a file
func FileMD5(file []byte) string {
	md5hash := md5.Sum(file)
	return hex.EncodeToString(md5hash[:])
}
