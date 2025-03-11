package numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratePrimes(t *testing.T) {
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59}
	actual := GeneratePrimes(60)

	assert.Equal(t, expected, actual)
}
