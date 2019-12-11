// Package circular implements a solution of the exercise titled `Circular Buffer'.
package circular

import "errors"

// Buffer is a struct of a circular buffer.
type Buffer struct {
	buffer      []byte
	top, bottom int
}

// NewBuffer is the ctor of Buffer.
func NewBuffer(size int) *Buffer {
	return &Buffer{buffer: make([]byte, size)}
}

// ReadByte returns a byte from the buffer or an error if buffer is empty.
func (b *Buffer) ReadByte() (byte, error) {
	if b.bottom+1 > b.top {
		return 0, errors.New("buffer is empty")
	}
	c := b.buffer[b.bottom%len(b.buffer)]
	b.bottom++
	return c, nil
}

// WriteByte stores a data or signals if buffer is full.
func (b *Buffer) WriteByte(c byte) error {
	if b.top-b.bottom+1 > len(b.buffer) {
		return errors.New("buffer is full")
	}
	b.buffer[b.top%len(b.buffer)] = c
	b.top++
	return nil
}

// Overwrite stores a data, overwrite the oldest data in the buffer when is's full.
func (b *Buffer) Overwrite(c byte) {
	b.buffer[b.top%len(b.buffer)] = c
	b.top++
	if b.top-b.bottom > len(b.buffer) {
		b.bottom++
	}
}

// Reset puts the buffer in empty state.
func (b *Buffer) Reset() {
	b.top, b.bottom = 0, 0
}

/*
BenchmarkOverwrite-4   	107902411	        11.1 ns/op	9744618248.23 MB/s	       0 B/op	       0 allocs/op
BenchmarkWriteRead-4   	54693124	        21.4 ns/op	2554800572.02 MB/s	       0 B/op	       0 allocs/op
*/
