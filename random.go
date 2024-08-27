package tripismapiutilities

import (
	"math/rand"
)

// RandomKeyCharacters is a []byte of the characters to choose from when generating random keys
var RandomKeyCharacters = []byte("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandomKey generates a random key at the given length
func RandomKey(length int) string {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		randInt := randInt(0, len(RandomKeyCharacters))
		bytes[i] = RandomKeyCharacters[randInt : randInt+1][0]
	}
	return string(bytes)
}

// randInt generates a random integer between min and max
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
