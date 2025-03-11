package numbers

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSqrt(t *testing.T) {
	assert.Equal(t, 12, Sqrt(144))
	assert.Equal(t, 7, Sqrt(49))
	n := rand.Intn(10000)
	assert.Equal(t, n, Sqrt(n*n))
}
