package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	v, ok := s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 4, v)

	v, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	v, ok = s.Peek()
	assert.True(t, ok)
	assert.Equal(t, 2, v)

	v, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 2, v)

	v, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	v, ok = s.Pop()
	assert.False(t, ok)
	assert.Equal(t, 0, v)
}
