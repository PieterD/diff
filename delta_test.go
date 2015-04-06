package diff_test

import (
	"testing"

	"github.com/PieterD/diff"
)

func TestDeltaString(t *testing.T) {
	if diff.Left.String() != "Left" {
		t.Fatalf("Expected Left, got %s", diff.Left.String())
	}
	if diff.Right.String() != "Right" {
		t.Fatalf("Expected Right, got %s", diff.Right.String())
	}
	if diff.Both.String() != "Both" {
		t.Fatalf("Expected Both, got %s", diff.Both.String())
	}
	if diff.Delta(55).String() != "unknown" {
		t.Fatalf("Expected unknown, got %s", diff.Delta(55).String())
	}
}
