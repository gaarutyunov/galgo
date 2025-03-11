package ringbuffer

import (
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	buffer := New[byte](make([]byte, 8))

	s := "hello"
	n, err := io.WriteString(buffer, s)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 5, n)

	b := make([]byte, n)
	n, err = io.ReadFull(buffer, b)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 5, n)
	assert.Equal(t, []byte("hello"), b[:n])
}

func TestRingBuffer_Overwrite(t *testing.T) {
	buffer := New[byte](make([]byte, 8))

	s := "hello, world!"

	n, err := io.WriteString(buffer, s)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(s), n)

	b := make([]byte, 8)
	n, err = io.ReadFull(buffer, b)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 8, n)
	assert.Equal(t, []byte(" world!,"), b)
}

func TestRingBuffer_EOF(t *testing.T) {
	buffer := New[byte](make([]byte, 8))

	s := "hello"
	n, err := io.WriteString(buffer, s)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 5, n)

	b := make([]byte, 8)
	n, err = io.ReadFull(buffer, b)
	assert.Equal(t, 5, n)
	assert.ErrorIs(t, io.ErrUnexpectedEOF, err)
}
