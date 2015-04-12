package diff_test

import (
	"fmt"

	"github.com/PieterD/diff"
)

func ExampleNew() {
	// Data
	l := []string{
		"linea",
		"lineb",
		"linec",
	}
	r := []string{
		"linea",
		"lineQ",
		"linec",
	}
	// Diff l and r using Strings
	d := diff.New(diff.Strings{
		Left:  l,
		Right: r,
	})
	// Print the diff
	for i := range d {
		switch d[i].Delta {
		case diff.Both:
			fmt.Printf("  %s\n", l[d[i].Index])
		case diff.Left:
			fmt.Printf("- %s\n", l[d[i].Index])
		case diff.Right:
			fmt.Printf("+ %s\n", r[d[i].Index])
		}
	}
	// Output:
	//   linea
	// - lineb
	// + lineQ
	//   linec
}
