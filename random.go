package tripismapiutilities

import (
	"math/rand"
	"sync"
	"time"
)

// RandomKeyCharacters is a []byte of the characters to choose from when generating random keys
var RandomKeyCharacters = []byte("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Use of Mutex to ensure function called once and only once
//
// See https://golang.org/src/sync/once.go
var randomOnce sync.Once

// RandomKey generates a random key at the given length
//
// The first time this is called, the rand.Seed will be set
// to the current time
func RandomKey(length int) string {
	// randomise the seed
	randomOnce.Do(func() {
		rand.Seed(time.Now().UTC().UnixNano())
	})
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
