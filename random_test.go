package tripismapiutilities

import (
	"testing"

	"github.com/cheekybits/is"
)

func TestRandom(t *testing.T) {
	is := is.New(t)
	is.Equal(32, len(RandomKey(32)))
}
