package diff_test

import (
	"testing"

	"github.com/PieterD/diff"
)

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

func bytesFromString(l, r string) Bytes {
	return Bytes{
		Left:  []byte(l),
		Right: []byte(r),
	}
}

type testStruct struct {
	t *testing.T
	d []diff.Diff
	p int
}

func newTestStruct(t *testing.T, d []diff.Diff, size int) testStruct {
	if len(d) != size {
		t.Fatalf("Wrong size: expected %d, got %d", size, len(d))
	}
	ts := testStruct{t, d, 0}
	return ts
}

func (ts *testStruct) expect(delta diff.Delta, index int) {
	if ts.p >= len(ts.d) {
		ts.t.Fatalf("Unexpected end")
	}
	if ts.d[ts.p].Delta != delta {
		ts.t.Fatalf("Wrong Delta idx %d: expected %d, got %d", ts.p, delta, ts.d[ts.p].Delta)
	}
	if ts.d[ts.p].Index != index {
		ts.t.Fatalf("Wrong Index idx %d: expected %d, got %d", ts.p, index, ts.d[ts.p].Index)
	}
	ts.p++
}

func (ts *testStruct) end() {
	if ts.p < len(ts.d) {
		ts.t.Fatalf("Expected end")
	}
}

func TestSimple(t *testing.T) {
	d := diff.New(bytesFromString("abc", "b"))
	ts := newTestStruct(t, d, 3)
	ts.expect(diff.Left, 0)
	ts.expect(diff.Both, 1)
	ts.expect(diff.Left, 2)
	ts.end()
}

func TestSimple2(t *testing.T) {
	d := diff.New(bytesFromString("abc", "bcde"))
	ts := newTestStruct(t, d, 5)
	ts.expect(diff.Left, 0)
	ts.expect(diff.Both, 1)
	ts.expect(diff.Both, 2)
	ts.expect(diff.Right, 2)
	ts.expect(diff.Right, 3)
	ts.end()
}

func TestSimple3(t *testing.T) {
	d := diff.New(bytesFromString("abcxyz", "bcdeyz"))
	ts := newTestStruct(t, d, 8)
	ts.expect(diff.Left, 0)
	ts.expect(diff.Both, 1)
	ts.expect(diff.Both, 2)
	ts.expect(diff.Left, 3)
	ts.expect(diff.Right, 2)
	ts.expect(diff.Right, 3)
	ts.expect(diff.Both, 4)
	ts.expect(diff.Both, 5)
	ts.end()
}

func TestLeft(t *testing.T) {
	d := diff.New(bytesFromString("aaaaaaa", "aaaaaaaaab"))
	ts := newTestStruct(t, d, 10)
	ts.expect(diff.Right, 0)
	ts.expect(diff.Right, 1)
	ts.expect(diff.Both, 0)
	ts.expect(diff.Both, 1)
	ts.expect(diff.Both, 2)
	ts.expect(diff.Both, 3)
	ts.expect(diff.Both, 4)
	ts.expect(diff.Both, 5)
	ts.expect(diff.Both, 6)
	ts.expect(diff.Right, 9)
}

func TestRight(t *testing.T) {
	d := diff.New(bytesFromString("aaaaaaaaab", "aaaaaaa"))
	ts := newTestStruct(t, d, 10)
	ts.expect(diff.Left, 0)
	ts.expect(diff.Left, 1)
	ts.expect(diff.Both, 2)
	ts.expect(diff.Both, 3)
	ts.expect(diff.Both, 4)
	ts.expect(diff.Both, 5)
	ts.expect(diff.Both, 6)
	ts.expect(diff.Both, 7)
	ts.expect(diff.Both, 8)
	ts.expect(diff.Left, 9)
}
