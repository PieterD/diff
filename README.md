diff
====

[![Coverage](http://gocover.io/_badge/github.com/PieterD/diff)](http://gocover.io/github.com/PieterD/diff)
[![GoDoc](https://godoc.org/github.com/PieterD/diff?status.svg)](https://godoc.org/github.com/PieterD/diff)

A small library to diff data.

It will work on any kind of data, as long as it can be expressed using the following interface:

    type Interface interface {
        // Return true if the elements at the given indices in Left and Right are equal.
        Equal(left, right int) (isEqual bool)
        // Return the sizes of the Left and Right collections.
        Length() (left int, right int)
    }

Pre-made implementations for strings and byte slices already exist. A simple example:

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
