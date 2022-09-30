package geecache

// ByteView holds an immutable view of bytes
type ByteView struct {
	b []byte
}

// Len return ByteView's length
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice return a copy of ByteView as a byte slice
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String return a string of ByteView data
func (v ByteView) String() string {
	return string(v.b)
}

// cloneBytes copy a ByteView's copy
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
