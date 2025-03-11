package ringbuffer

import (
	"io"
)

type RingBuffer[T any] struct {
	buffer            []T
	writeIdx, readIdx int
}

func New[T any](buffer []T) *RingBuffer[T] {
	return &RingBuffer[T]{
		buffer: buffer,
	}
}

func (r *RingBuffer[T]) Empty() bool {
	return r.writeIdx == r.readIdx
}

func (r *RingBuffer[T]) Cap() int {
	return len(r.buffer)
}

func (r *RingBuffer[T]) Write(b []T) (n int, err error) {
	readIdx, writeIdx := r.readIdx, r.writeIdx
	defer func() {
		r.readIdx = readIdx
		r.writeIdx = writeIdx
	}()

	i := 0
	for ; i < len(b); i++ {
		r.buffer[writeIdx] = b[i]
		writeIdx = (writeIdx + 1) % len(r.buffer)
		if writeIdx == readIdx {
			readIdx = (readIdx + 1) % len(r.buffer)
		}
	}

	return i, nil
}

func (r *RingBuffer[T]) Read(b []T) (n int, err error) {
	readIdx, writeIdx := r.readIdx, r.writeIdx
	if readIdx == writeIdx {
		return 0, io.EOF
	}
	defer func() {
		r.readIdx = readIdx
	}()

	i := 0
	for ; i < len(b); i++ {
		b[i] = r.buffer[readIdx]

		if readIdx == writeIdx && i < len(b)-1 {
			return i, io.EOF
		}

		readIdx = (readIdx + 1) % len(r.buffer)
	}

	return i, nil
}
