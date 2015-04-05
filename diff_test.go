package diff_test

import (
	"testing"

	"github.com/PieterD/diff"
)

func bytesFromString(l, r string) diff.Bytes {
	return diff.Bytes{
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
		t.Fatalf("Wrong size: expected %d, got %d", len(d), size)
	}
	ts := testStruct{t, d, 0}
	return ts
}

func (ts *testStruct) expect(delta diff.Delta, index int) {
	if ts.p >= len(ts.d) {
		ts.t.Fatalf("Unexpected end")
	}
	if ts.d[ts.p].Delta != delta {
		ts.t.Fatalf("Wrong Delta idx %d: expected %d, got %d", ts.p, ts.d[ts.p].Delta, delta)
	}
	if ts.d[ts.p].Index != index {
		ts.t.Fatalf("Wrong Index idx %d: expected %d, got %d", ts.p, ts.d[ts.p].Index, index)
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
