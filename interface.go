package diff

// Wrap your data in one of these to diff it.
// It should hold two collections, the Left (or old) one and the Right (or new) one.
type Interface interface {
	// Return true if the elements at the given indices in Left and Right are equal.
	Equal(left, right int) (isEqual bool)
	// Return the sizes of the Left and Right collections.
	Length() (left int, right int)
}

type Delta int

const (
	// Element is present in both Left and Right collections.
	// Index uses the Left collection.
	Both Delta = iota
	// Element is present only in the Left collection.
	// Index uses the Left collection.
	Left
	// Element is present only in the Right collection.
	// Index uses the Right collection.
	Right
)

func (delta Delta) String() string {
	switch delta {
	case Both:
		return "Both"
	case Left:
		return "Left"
	case Right:
		return "Right"
	}
	return "unknown"
}

// One Diff record per element.
// If Delta is Left or Both, Index is for the left collection.
// If Delta is Right, Index is for the right collection.
type Diff struct {
	Delta Delta
	Index int
}

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

// Holds byte slices to diff, one character at a time.
type Bytes struct {
	Left, Right []byte
}

func (b Bytes) Equal(left, right int) bool {
	return b.Left[left] == b.Right[right]
}

func (b Bytes) Length() (int, int) {
	return len(b.Left), len(b.Right)
}
