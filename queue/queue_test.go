package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	q := New[int]()
	q.Push(4)
	q.Push(3)
	q.Push(2)
	q.Push(1)

	v, ok := q.Pop()
	assert.True(t, ok)
	assert.Equal(t, 4, v)

	v, ok = q.Pop()
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	v, ok = q.Peek()
	assert.True(t, ok)
	assert.Equal(t, 2, v)

	v, ok = q.Pop()
	assert.True(t, ok)
	assert.Equal(t, 2, v)

	v, ok = q.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	v, ok = q.Pop()
	assert.False(t, ok)
	assert.Equal(t, 0, v)
}
