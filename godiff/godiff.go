package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/PieterD/diff"
)

var (
	plus  = []byte("+ ")
	minus = []byte("- ")
	blank = []byte("  ")
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: godiff <old> <new>\n")
		os.Exit(1)
	}
	left := read(args[0])
	right := read(args[1])
	d := diff.New(diff.Bytes{
		Left:  left,
		Right: right,
	})

	for i := range d {
		var pfx, str []byte
		switch d[i].Delta {
		case diff.Both:
			pfx = blank
			str = left[d[i].Index]
		case diff.Left:
			pfx = minus
			str = left[d[i].Index]
		case diff.Right:
			pfx = plus
			str = right[d[i].Index]
		}
		fmt.Printf("%s%s\n", pfx, str)
	}
}

func read(filename string) [][]byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return bytes.Split(b, []byte("\n"))
}
