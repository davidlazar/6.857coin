package coin

import (
	"testing"
)

var tests = []struct {
	hash       string
	difficulty int
	ok         bool
}{
	{"ff00000000000000000000000000000000000000000000000000000000000000", 0, true},
	{"ff00000000000000000000000000000000000000000000000000000000000000", 1, false},
	{"ff00000000000000000000000000000000000000000000000000000000000000", 8, false},
	{"ff00000000000000000000000000000000000000000000000000000000000000", 9, false},
	{"0f00000000000000000000000000000000000000000000000000000000000000", 4, true},
	{"0f00000000000000000000000000000000000000000000000000000000000000", 5, false},
	{"0010000000000000000000000000000000000000000000000000000000000000", 11, true},
	{"0010000000000000000000000000000000000000000000000000000000000000", 12, false},
	{"0000000000000000000000000000000000000000000000000000000000000000", 256, true},
	{"0000000000000000000000000000000000000000000000000000000000000001", 256, false},
	{"0000000000000000000000000000000000000000000000000000000000000001", 255, true},
	{"000007e5a9ff1e4464e059265bd6f21170b8342296100d7a847abca7675f54ba", 21, true},
	{"000007e5a9ff1e4464e059265bd6f21170b8342296100d7a847abca7675f54ba", 22, false},
}

func TestSlowCheck(t *testing.T) {
	for i, test := range tests {
		h, _ := NewHash(test.hash)
		if slowCheck(test.difficulty, &h) != test.ok {
			t.Fatalf("failed test %d", i)
		}
	}
}

func TestFastCheck(t *testing.T) {
	for i, test := range tests {
		h, _ := NewHash(test.hash)
		if fastCheck(test.difficulty, &h) != test.ok {
			t.Fatalf("failed test %d", i)
		}
	}
}

func BenchmarkSlowCheck(b *testing.B) {
	t := tests[len(tests)-1]
	h, _ := NewHash(t.hash)
	for i := 0; i < b.N; i++ {
		_ = slowCheck(t.difficulty, &h)
	}
}

func BenchmarkFastCheck(b *testing.B) {
	t := tests[len(tests)-1]
	h, _ := NewHash(t.hash)
	for i := 0; i < b.N; i++ {
		_ = fastCheck(t.difficulty, &h)
	}
}
