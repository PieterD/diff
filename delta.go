package diff

// Describe in which collection the element occurs; Left, Right or Both.
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
