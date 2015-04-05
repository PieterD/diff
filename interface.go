package diff

type Interface interface {
	Equal(left, right int) (isEqual bool)
	Length() (left int, right int)
}

type Delta int

const (
	Both Delta = iota
	Left
	Right
)

// One Diff record per element.
// If Delta is Left or Both, Index is for the left collection.
// If Delta is Right, Index is for the right collection.
type Diff struct {
	Delta Delta
	Index int
}

type Strings struct {
	Left, Right []string
}

func (str Strings) Equal(left, right int) bool {
	return str.Left[left] == str.Right[right]
}

func (str Strings) Length() (int, int) {
	return len(str.Left), len(str.Right)
}

type Bytes struct {
	Left, Right []byte
}

func (b Bytes) Equal(left, right int) bool {
	return b.Left[left] == b.Right[right]
}

func (b Bytes) Length() (int, int) {
	return len(b.Left), len(b.Right)
}
