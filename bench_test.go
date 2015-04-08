package diff_test

import (
	"math/rand"
	"testing"

	"github.com/PieterD/diff"
)

func genStr(size int, ratio float64) []byte {
	b := make([]byte, size)
	for i := range b {
		c := byte('a')
		if rand.Float64() < ratio {
			c = 'b'
		}
		b[i] = c
	}
	return b
}

var str = genStr(500, 0.5)

func strPart(b []byte) []byte {
	i := rand.Int31n(int32(len(b)))
	j := i + rand.Int31n(int32(len(b))-i)
	return b[i:j]
}

func BenchmarkSame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		diff.New(Bytes{[]byte(str), []byte(str)})
	}
}

func BenchmarkNotRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		diff.New(Bytes{[]byte("ababaaabababbababbaabbaabababaa"), []byte("aaababbbababaabababababbabababababababababababbabbbabababaaabbb")})
	}
}

func BenchmarkRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := strPart(str)
		b := strPart(str)
		diff.New(Bytes{a, b})
	}
}

func BenchmarkLeft(b *testing.B) {
	for i := 0; i < b.N; i++ {
		left := str[250:]
		right := str
		diff.New(Bytes{left, right})
	}
}

func BenchmarkRight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		left := str
		right := str[250:]
		diff.New(Bytes{left, right})
	}
}
