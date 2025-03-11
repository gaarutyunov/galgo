package hashmap

import (
	"github.com/gaarutyunov/galgo/functional"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashMap_Strings(t *testing.T) {
	m := New[string, int](3, RabinFingerprint)

	m.Set("Alice", 4)
	m.Set("Bob", 3)
	m.Set("Tom", 5)

	v, ok := m.Get("Alice")
	assert.True(t, ok)
	assert.Equal(t, 4, v)

	v, ok = m.Get("Bob")
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	v, ok = m.Get("Tom")
	assert.True(t, ok)
	assert.Equal(t, 5, v)

	assert.Equal(t, 3, m.Len())

	m.Delete("Alice")

	v, ok = m.Get("Alice")
	assert.False(t, ok)

	v, ok = m.Get("Bob")
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	v, ok = m.Get("Tom")
	assert.True(t, ok)
	assert.Equal(t, 5, v)

	assert.Equal(t, 2, m.Len())

	v, ok = m.Get("Glen")
	assert.False(t, ok)

	m.Set("Bob", 4)

	v, ok = m.Get("Bob")
	assert.True(t, ok)
	assert.Equal(t, 4, v)

	assert.Equal(t, 2, m.Len())

	m.Set("Alice", 4)

	assert.Equal(t, 3, m.Len())

	m.Set("Alice", 2)

	assert.Equal(t, 3, m.Len())

	m.Delete("Alice")

	assert.Equal(t, 2, m.Len())

	m.Set("Alice", 2)

	assert.Equal(t, 3, m.Len())

	m.Delete("Bob")

	assert.Equal(t, 2, m.Len())
}

func TestHashMap_Int(t *testing.T) {
	m := New[int, struct{}](10, Int)

	m.Set(1, struct{}{})
	m.Set(2, struct{}{})

	v, ok := m.Get(1)
	assert.True(t, ok)
	assert.Equal(t, struct{}{}, v)
	assert.Equal(t, 2, m.Len())
}

func TestHashMap_DJBA2(t *testing.T) {
	m := New[string, int](10, functional.Pipe2(functional.StringToByteSlice, DJBA2))

	m.Set("test", 1)
	m.Set("test2", 2)

	assert.Equal(t, 2, m.Len())
}
