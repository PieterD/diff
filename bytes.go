package diff

import "bytes"

// Holds two string lists to diff.
type Bytes struct {
	Left, Right [][]byte
}

func (str Bytes) Equal(left, right int) bool {
	return bytes.Equal(str.Left[left], str.Right[right])
}

func (str Bytes) Length() (int, int) {
	return len(str.Left), len(str.Right)
}
