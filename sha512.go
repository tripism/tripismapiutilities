package tripismapiutilities

import (
	"crypto/sha512"
	"fmt"
)

// CalculateSHA512 generates a string representation of the SHA512 hash for a input string
func CalculateSHA512(s string) string {
	if len(s) == 0 {
		return ""
	}
	hasher := sha512.New()
	hasher.Write([]byte(s))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
