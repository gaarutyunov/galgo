package hashmap

import (
	"github.com/gaarutyunov/galgo/numbers"
	"unsafe"
)

const alphabetSize = 1 << 7

func RabinFingerprint(s string) (hash uint64) {
	for i, j := 0, len(s)-1; i < len(s); i, j = i+1, j-1 {
		hash += uint64(s[j]) * (alphabetSize << i)
	}

	return
}

func Ptr[K comparable](v K) (hash uint64) {
	return uint64(uintptr(unsafe.Pointer(&v)))
}

func Int[T numbers.Integer](n T) (hash uint64) {
	return uint64(n)
}

func DJBA2(data []byte) uint64 {
	hash := uint64(5381)

	for _, b := range data {
		hash += uint64(b) + hash + hash<<5
	}

	return hash
}
