package diff

// Holds two string lists to diff.
type Strings struct {
	Left, Right []string
}

func (str Strings) Equal(left, right int) bool {
	return str.Left[left] == str.Right[right]
}

func (str Strings) Length() (int, int) {
	return len(str.Left), len(str.Right)
}
