package tripismapiutilities

import (
	"crypto/rand"
)

// RandomKeyCharacters is a []byte of the characters to choose from when generating random keys
var RandomKeyCharacters = []byte("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandomKey generates a random key at the given length
func RandomKey(length int) string {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		b := make([]byte, 1)
		_, err := rand.Read(b)
		if err != nil {
			return ""
		}
		bytes[i] = RandomKeyCharacters[int(b[0])%len(RandomKeyCharacters)]
	}
	return string(bytes)
}
